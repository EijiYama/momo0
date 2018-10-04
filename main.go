package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"log"
	"sync"
	"time"

	"github.com/momo0/test001/uc"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

const (
	tickDuration = time.Second / 60

	defWinSX = 800
	defWinSY = 600

	pause = false
	play  = true
)

var Hdl = struct {
	mu              sync.Mutex
	uploadEventSent bool
	mouseEvents     []image.Point
}{}

var pauseChan = make(chan bool, 64)

type uploadEvent struct{}

func main() {
	driver.Main(func(s screen.Screen) {

		w, err := s.NewWindow(&screen.NewWindowOptions{Title: "Test001", Width: defWinSX, Height: defWinSY})
		if err != nil {
			log.Fatal(err)
		}
		buf, tex := screen.Buffer(nil), screen.Texture(nil)
		defer func() {
			if buf != nil {
				tex.Release()
				buf.Release()
			}
			w.Release()
		}()

		go simulate(w)

		var (
			buttonDown bool
			sz         size.Event
		)
		for {
			publish := false

			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				fmt.Println("lifecycle.Event")
				if e.To == lifecycle.StageDead {
					return
				}

				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					fmt.Println("lifecycle.CrossOn App.WinSX, App.WinSY :", uc.App.WinSX, uc.App.WinSY)
					pauseChan <- play
					if buf, err = s.NewBuffer(image.Point{uc.App.WinSX, uc.App.WinSY}); err != nil {
						log.Fatal(err)
					}
					if tex, err = s.NewTexture(image.Point{uc.App.WinSX, uc.App.WinSY}); err != nil {
						log.Fatal(err)
					}
					tex.Fill(tex.Bounds(), color.White, draw.Src)

				case lifecycle.CrossOff:
					fmt.Println("lifecycle.CrossOff")
					pauseChan <- pause
					tex.Release()
					tex = nil
					buf.Release()
					buf = nil
				}

			case mouse.Event:
				if e.Button == mouse.ButtonLeft {
					buttonDown = e.Direction == mouse.DirPress
				}
				uc.App.Ms.Px, uc.App.Ms.Py = int(e.X), int(e.Y)
				//				fmt.Println("mouse.Event (x, y)", int(e.X), int(e.Y))
				if !buttonDown {
					break
				}
				/*				z := sz.Size()
								x, y := int(e.X) *N / z.X, int(e.Y) *N / z.Y   //*/
				x, y := int(e.X), int(e.Y) //
				//				fmt.Println("x, y, z :", x, y, z)
				if x < 0 || sz.WidthPx <= x || y < 0 || sz.HeightPx <= y {
					break
				}

				Hdl.mu.Lock()
				Hdl.mouseEvents = append(Hdl.mouseEvents, image.Point{x, y})
				Hdl.mu.Unlock()

			case paint.Event:
				publish = buf != nil

			case size.Event:
				sz = e

				//メイン画面初期化
				uc.InitApp(sz.WidthPx, sz.HeightPx)
				fmt.Println("size.Event  sz, sz.WidthPx, sz.HeightPx :", sz, sz.WidthPx, sz.HeightPx)

				if buf, err = s.NewBuffer(image.Point{uc.App.WinSX, uc.App.WinSY}); err != nil {
					log.Fatal(err)
				}
				if tex, err = s.NewTexture(image.Point{uc.App.WinSX, uc.App.WinSY}); err != nil {
					log.Fatal(err)
				}

			case uploadEvent:
				Hdl.mu.Lock()
				if buf != nil {
					copy(buf.RGBA().Pix, uc.App.Pix)
					publish = true
				}
				Hdl.uploadEventSent = false
				Hdl.mu.Unlock()

				if publish {
					tex.Upload(image.Point{}, buf, buf.Bounds())
				}

			case error:
				log.Print(e)
			}

			if publish {
				w.Scale(sz.Bounds(), tex, tex.Bounds(), draw.Src, nil)
				w.Publish()
			}
		}
	})
}

func simulate(q screen.EventDeque) {

	uc.InitApp(uc.App.WinSX, uc.App.WinSY)

	ticker := time.NewTicker(tickDuration)
	var tickerC <-chan time.Time
	for {
		select {
		case p := <-pauseChan:
			fmt.Println("pauseChan p :", p)
			if p == pause {
				tickerC = nil
			} else {
				tickerC = ticker.C
			}
			continue
		case <-tickerC:
		}

		Hdl.mu.Lock()

		for _, p := range Hdl.mouseEvents {
			fmt.Println("simulate mouse Evenet x, y :", p.X, p.Y)

			uc.MouseEvent()

		}
		Hdl.mouseEvents = Hdl.mouseEvents[:0]
		uc.UpdScn()
		uploadEventSent := Hdl.uploadEventSent
		Hdl.uploadEventSent = true
		Hdl.mu.Unlock()

		if !uploadEventSent {
			q.Send(uploadEvent{})
		}
	}
}
