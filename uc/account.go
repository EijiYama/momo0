package uc

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/momo0/test001/cn"
	"github.com/momo0/test001/tl"
	"golang.org/x/image/font"
)

func (w *Window) DrawAccEqbalFrame() {

	//表題エリア
	titleH := 20
	fmt.Println("DrawAccEqbalFrame sx, ex :", w.Fd.sx, w.Fd.ex)
	for y := w.Fd.sy; y < w.Fd.sy+titleH; y++ {
		tl.DrawLine(w.Win, cn.ColHHDarkGray, w.Fd.sx, y, w.Fd.ex, y, 0)
	}

	rf := NewFrame(w.Fd.sx, w.Fd.sy+titleH, w.Fd.ex, w.Fd.ey) //一覧領域着色
	for y := rf.sy; y < rf.ey; y++ {
		col := cn.ColGray33
		if (y-rf.sy)/lineH%2 == 1 {
			col = cn.ColGray22
		}
		tl.DrawLine(w.Win, col, rf.sx, y, rf.ex, y, 0)
	}
}

//現物残高
func (w *Window) DrawAccEqBalCont(eqbal *EquityBalancesResponse, dx int, dy int) {

	face := truetype.NewFace(App.Font, &truetype.Options{Size: fontS, DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawSimple(w.Win, App.accEqbalWin.Win, dx, dy)

	interval := 2 //表示間隔
	clmLen := []int{5, 14, 8, 8, 12, 8, 9, 9}
	clmStr := []string{"銘柄", "", "市場", "数量", "単価", "口座", "最終約定価格", "時価更新"}

	x := dx + w.Fd.sx + 5
	for i := 0; i < len(clmLen); i++ {
		tl.DrawText(w.Win, cn.ColDarkBlue, face, x, dy+w.Fd.sy+15, clmStr[i])
		x += clmLen[i]*int(fontS)/2 + interval
	}

	for j, eb := range eqbal.EquityBalances {
		tblStr := []string{tl.PS2S(eb.StockCode), tl.PS2S(eb.StockName), tl.PS2S(eb.MarketSection), tl.PS2S(eb.BalanceQuantity), tl.PS2S(eb.BookUnitPrice), tl.PS2S(eb.AccountType), tl.PS2S(eb.CurrentPrice), tl.PS2S(eb.CurrentPriceTime)}

		x = dx + w.Fd.sx + 5
		y := dy + w.Fd.sy + 15 + 2 + (j+1)*14

		for i := 0; i < len(clmLen); i++ {
			tl.DrawText(w.Win, cn.ColWhite, face, x, y, tblStr[i])
			x += clmLen[i]*int(fontS)/2 + interval
		}
	}
}
