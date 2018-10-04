package uc

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/momo0/test001/cn"
	"github.com/momo0/test001/tl"
)

var (
	btn = map[int]*Button{
		BOARD:    {dispRc: image.Rect(50, 0, 100, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_board_off.jpg"), imgOff: tl.LoadImg("btn_board_on.jpg"), pressed: false},
		CHART:    {dispRc: image.Rect(110, 0, 160, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_chart_off.jpg"), imgOff: tl.LoadImg("btn_chart_on.jpg"), pressed: false},
		MV_ASSET: {dispRc: image.Rect(170, 0, 220, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_board_off.jpg"), imgOff: tl.LoadImg("btn_board_on.jpg"), pressed: false},
		NEWS:     {dispRc: image.Rect(230, 0, 280, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_news_off.jpg"), imgOff: tl.LoadImg("btn_news_on.jpg"), pressed: false},
		TOSHI:    {dispRc: image.Rect(290, 0, 340, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_toshi_off.jpg"), imgOff: tl.LoadImg("btn_toshi_on.jpg"), pressed: false},
		NOTICE:   {dispRc: image.Rect(350, 0, 400, 50), imgRc: image.Rect(0, 0, 50, 50), imgOn: tl.LoadImg("btn_notice_off.jpg"), imgOff: tl.LoadImg("btn_notice_on.jpg"), pressed: false},
	}
	imgOnMouse = tl.LoadImg("btn_on_mouse.jpg")
	imgSel     = tl.LoadImg("btn_selected.jpg")
)

type Button struct {
	pressed bool
	imgRc   image.Rectangle
	dispRc  image.Rectangle
	imgOn   *image.RGBA
	imgOff  *image.RGBA
}

func (w *Window) DrawMMenuFrame() {

	//	logo := tl.LoadImg("smartplus_log0.jpg")
	//	draw.Draw(w.Win, image.Rect(2, 2, 2+logo.Bounds().Dx(), 2+logo.Bounds().Dy()), logo, cn.ZP, draw.Src) //ロゴ

	/*	   mainMenu := tl.LoadImg("main_menu.jpg")	*/

	for _, b := range btn {
		draw.Draw(w.Win, b.dispRc, b.imgOff, cn.ZP, draw.Src)
	}

	for x := 0; x < imgSel.Bounds().Dx(); x++ { //メインメニューボタン押下用のα値設定
		for y := 0; y < imgSel.Bounds().Dy(); y++ {

			r, g, b, _ := imgSel.At(x, y).RGBA()
			imgSel.SetRGBA(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 0xa0})

			r, g, b, _ = imgOnMouse.At(x, y).RGBA()
			imgOnMouse.SetRGBA(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 0xa0})
		}
	}
}

func (w *Window) DrawMMenuCont(dx int, dy int) {

	//	face := truetype.NewFace(App.Font, &truetype.Options{Size: fontS, DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawSimple(w.Win, App.mMenu.Win, App.mMenu.X, App.mMenu.Y)

	mx, my := App.Ms.Px, App.Ms.Py
	for i, b := range btn {

		draw.Draw(w.Win, b.dispRc, b.imgOff, cn.ZP, draw.Src)

		if dx+b.dispRc.Bounds().Min.X < mx && mx < dx+b.dispRc.Bounds().Max.X &&
			dy+b.dispRc.Bounds().Min.Y < my && my < dy+b.dispRc.Bounds().Max.Y { //カーソルON
			draw.DrawMask(w.Win, b.dispRc, imgOnMouse, cn.ZP, imgOnMouse, cn.ZP, draw.Over)
		}

		if App.selMMenu == i { //メニュー選択
			draw.DrawMask(w.Win, b.dispRc, imgSel, cn.ZP, imgSel, cn.ZP, draw.Over)
		}

	}
}
