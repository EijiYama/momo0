package tl

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func DrawText(m *image.RGBA, fcol color.RGBA, fnt font.Face, x int, y int, text string) {
	//	draw.Draw(m, m.Bounds(), image.White, ZP, draw.Src)	//文字の背景塗りつぶし（ここではやらない）
	//	m.SetRGBA(x, y, crossColor)

	d := &font.Drawer{
		Dst:  m,
		Src:  image.NewUniform(fcol),
		Face: fnt, //example : inconsolata.Regular8x16, basicfont.Face7x13
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(y * 64),
		},
	}

	d.DrawString(text)

}

func LoadImg(name string) *image.RGBA {
	file, err := os.Open(filepath.Join("img", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	imgRGB := image.NewRGBA(img.Bounds())
	draw.Draw(imgRGB, imgRGB.Bounds(), img, image.Point{}, draw.Src)

	return imgRGB
}

func LoadFont(path string) *truetype.Font {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fontBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Fatal(err2)
	}

	tmp, err3 := truetype.Parse(fontBytes)
	if err3 != nil {
		log.Fatal(err3)
	}
	return tmp
}

func DrawLine(m *image.RGBA, fcol color.RGBA, sx, sy, ex, ey int, grdTerm int) {

	if sx > ex {
		tmp := sx
		sx = ex
		ex = tmp
	}

	if sy > ey {
		tmp := sy
		sy = ey
		ey = tmp
	}

	if sx == ex {
		for y := sy; y <= ey; y++ {
			if grdTerm == 0 || (y/grdTerm%2) == 0 {
				m.SetRGBA(sx, y, fcol)
			}
		}
	} else if sy == ey {
		for x := sx; x <= ex; x++ {
			if grdTerm == 0 || (x/grdTerm%2) == 0 {
				m.SetRGBA(x, sy, fcol)
			}
		}
	} else {

		wx := ex - sx
		wy := ey - sy
		D := 2*wy - wx
		m.SetRGBA(sx, sy, fcol)
		y := sy

		for x := sx + 1; x < ex; x++ {
			if D > 0 {
				y++
				if grdTerm == 0 || (x/grdTerm%2) == 0 {
					m.SetRGBA(x, y, fcol)
				}
				D = D + (2*wy - 2*wx)
			} else {
				if grdTerm == 0 || (x/grdTerm%2) == 0 {
					m.SetRGBA(x, y, fcol)
				}
				D = D + (2 * wy)
			}
		}
	}
}

func DrawFrame(m *image.RGBA, fcol color.RGBA, sx, sy, ex, ey int, grdTerm int) {
	DrawLine(m, fcol, sx, sy, sx, ey, grdTerm) //左線
	DrawLine(m, fcol, ex, sy, ex, ey, grdTerm) //右線
	DrawLine(m, fcol, sx, sy, ex, sy, grdTerm) //上線
	DrawLine(m, fcol, sx, ey, ex, ey, grdTerm) //下線
}

func FillFrame(m *image.RGBA, fcol *color.RGBA, sx, sy, ex, ey int) {

	for x := sx; x < ex; x++ {
		for y := sy; y < ey; y++ {
			m.SetRGBA(x, y, *fcol)
		}
	}
}

func FillRect(m *image.RGBA, fcol *color.RGBA, rc *image.Rectangle) {

	for x := rc.Bounds().Min.X; x < rc.Bounds().Max.X; x++ {
		for y := rc.Bounds().Min.Y; y < rc.Bounds().Max.Y; y++ {
			m.SetRGBA(x, y, *fcol)
		}
	}
}

func DrawSimple(dst *image.RGBA, src *image.RGBA, x int, y int) {
	draw.Draw(dst, image.Rect(x, y, x+src.Bounds().Dx(), y+src.Bounds().Dy()), src, image.ZP, draw.Src)
}
