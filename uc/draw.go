package uc

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/momo0/test001/cn"
	"github.com/momo0/test001/tl"
	"golang.org/x/image/font"
)

type Frame struct {
	sx, sy, ex, ey int //開始点(x, y)、終了点(x, y)
	w, h           int //幅、高さ
}

func NewFrame(sx, sy, ex, ey int) *Frame {
	f := Frame{sx: sx, sy: sy, ex: ex, ey: ey}
	f.w = f.ex - f.sx
	f.h = f.ey - f.sy

	return &f
}

type Window struct {
	Win    *image.RGBA //型枠
	X, Y   int         //表示位置（メインウインドウでの相対位置）
	W, H   int         //幅、高さ
	Rc     *image.Rectangle
	BgCol  *color.RGBA
	FntCol *color.RGBA
	Type   string
	Fd     *Frame //表示領域フレーム
}

//Windowリサイズによる際呼び出しに対応
func InitApp(winSX int, winSY int) {
	App.selMMenu = -1
	App.WinSX, App.WinSY = winSX, winSY
	App.Pix = make([]uint8, App.WinSX*App.WinSY*4)
	App.Font = tl.LoadFont("font/msgothic.ttc") //	App.Font := tl.LoadFont("font/HGRSMP.TTF")

	App.baseScn = NewWindow(&image.Rectangle{cn.ZP, image.Point{App.WinSX, App.WinSY}}, &cn.ColBlack, &cn.ColWhite, cn.WINTYPE_NOFRAME) //メイン画面ひな形
	App.outScn = NewWindow(&image.Rectangle{cn.ZP, image.Point{App.WinSX, App.WinSY}}, &cn.ColBlack, &cn.ColWhite, cn.WINTYPE_NOFRAME)  //メイン画面表示用領域

	App.mMenu = NewWindow(&image.Rectangle{cn.ZP, image.Point{App.WinSX, 70}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_NOFRAME) //メイン画面ひな形
	App.mMenu.DrawMMenuFrame()

	App.accEqbalWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{600, 250}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //口座残高 株式残高
	App.accEqbalWin.DrawAccEqbalFrame()

	//投資情報　左列
	App.invAssetWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{200, 130}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY)
	App.invAssetWin.DrawInvAssetFrame()
	App.invEqbalWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{200, 300}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //株式残高
	App.invEqbalWin.DrawInvEqbalFrame()
	App.invMktWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{200, 100}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //マーケット
	App.invMktWin.DrawInvMarketFrame()
	//投資情報　中央列
	App.invPriceWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{200, 330}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //投資情報　板情報
	App.invPriceWin.DrawInvPriceFrame()
	App.invOrderWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{200, 600}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //投資情報　注文
	App.invOrderWin.DrawInvOrderFrame()
	//投資情報　右列
	App.invBalDtlWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{400, 130}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //投資情報　銘柄残高詳細
	App.invBalDtlWin.DrawInvBalDtlFrame()
	App.invCBarWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{400, 200}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //投資情報　チャートバー
	App.invCBarWin.DrawInvCBarFrame()
	App.invCDtlWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{400, 400}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //投資情報　チャート詳細
	App.invCDtlWin.DrawInvCDtlFrame()

	App.chartWin = NewWindow(&image.Rectangle{cn.ZP, image.Point{800, 530}}, &cn.ColHDarkGray, &cn.ColWhite, cn.WINTYPE_FRAMEIONLY) //チャート
	App.chartWin.DrawChartFrame()

	App.updTime = time.Now()
}

func UpdScn() {
	nowTime := time.Now()

	//		draw.Draw(App.outScn.Win, *App.baseScn.Rc, App.baseScn.Win, ZP, draw.Src)
	copy(App.outScn.Win.Pix, App.baseScn.Win.Pix)

	App.outScn.DrawMMenuCont(0, 0) //メインメニュー(TODO 高速化可能＜outScnでほとんどを書く)

	switch App.selMMenu {
	case cn.MENU_BOARD:
	case cn.MENU_MV_ASSET:
	case cn.MENU_CHART:
		App.outScn.DrawChartCont(0, 70) //チャート
	case cn.MENU_NEWS:
	case cn.MENU_INVEST:
		App.outScn.DrawInvAssetCont(0, 70)
		App.outScn.DrawInvEqBalCont(GetTestDataEquityBalances(), 0, 200) //株式残高
		App.outScn.DrawInvMktCont(0, 500)
		App.outScn.DrawInvPriceCont(200, 70)
		App.outScn.DrawInvOrderCont(200, 400)
		App.outScn.DrawInvBalDtlCont(400, 70)
		App.outScn.DrawInvCBarCont(400, 200)
		App.outScn.DrawInvCDtlCont(400, 400)
	case cn.MENU_ACCOUNT:
		App.outScn.DrawAccEqBalCont(GetTestDataEquityBalances(), 0, 70) //株式残高
	case cn.MENU_NOTICE:
	}

	App.outScn.DrawSystemInfo(nowTime)

	copy(App.Pix, App.outScn.Win.Pix)
	App.updTime = nowTime
}

//システム情報
func (w *Window) DrawSystemInfo(nowTime time.Time) {

	fntFace12 := truetype.NewFace(App.Font, &truetype.Options{Size: 12, DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawText(w.Win, cn.ColWhite, fntFace12, 500, 10, fmt.Sprintf("now :%+v", nowTime.Format("2006/1/2 15:04:05.999999")))
	tl.DrawText(w.Win, cn.ColWhite, fntFace12, 500, 20, fmt.Sprintf("frame time :%+v", nowTime.Sub(App.updTime)))
	if nowTime.Sub(App.updTime).Nanoseconds() > 0 {
		tl.DrawText(w.Win, color.RGBA{255, 100, 100, 255}, fntFace12, 650, 20, fmt.Sprintf("fps :%-3.1f", ((float64)(1000000000.0)/(float64)(nowTime.Sub(App.updTime).Nanoseconds()))))
	}
}

func NewWindow(winrc *image.Rectangle, bgCol *color.RGBA, fntCol *color.RGBA, winType string) *Window {

	w := &Window{
		Rc: winrc, W: winrc.Bounds().Dx(), H: winrc.Bounds().Dy(),
		Win: image.NewRGBA(*winrc), BgCol: bgCol, FntCol: fntCol, Type: winType,
	}

	tl.FillRect(w.Win, w.BgCol, w.Rc)
	w.DrawWinFrame()

	return w
}

func (w *Window) DrawWinFrame() {

	wx, wy := w.Rc.Dx(), w.Rc.Dy()

	//実表示領域設定
	fWidth := 2

	w.Fd = NewFrame(fWidth, fWidth, wx-fWidth, wy-fWidth)

	if w.Type == cn.WINTYPE_TYPE1 || w.Type == cn.WINTYPE_FRAMEIONLY { //フレーム描画(対象外：WINTYPE_NOFRAME)
		tl.DrawLine(w.Win, cn.ColHDarkGray, 0, 0, wx, 0, 0)
		tl.DrawLine(w.Win, cn.ColDarkGray, 0, 1, wx, 1, 0)

		tl.DrawLine(w.Win, cn.ColHDarkGray, 0, wy-2, wx, wy-2, 0)
		tl.DrawLine(w.Win, cn.ColDarkGray, 0, wy-1, wx, wy-1, 0)

		tl.DrawLine(w.Win, cn.ColHDarkGray, 0, 0, 0, wy, 0)
		tl.DrawLine(w.Win, cn.ColDarkGray, 1, 1, 1, wy-1, 0)

		tl.DrawLine(w.Win, cn.ColHDarkGray, wx-2, 2, wx-2, wy-1, -1)
		tl.DrawLine(w.Win, cn.ColDarkGray, wx-1, 1, wx-1, wy, 0)
	}

	//タイトルバー描画
	if w.Type == cn.WINTYPE_TYPE1 {
		for y := 3; y < 10; y++ {
			tl.DrawLine(w.Win, cn.ColLightGray, 2, y, wx-3, y, 0)
		}
	} else if w.Type == cn.WINTYPE_TYPE2 {
		for y := 3; y < 20; y++ {
			tl.DrawLine(w.Win, cn.ColLightGray, 2, y, wx-3, y, 0)
		}
	}

}
