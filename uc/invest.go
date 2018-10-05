package uc

import (
	"github.com/golang/freetype/truetype"
	"github.com/momo0/test001/cn"
	"github.com/momo0/test001/tl"
	"golang.org/x/image/font"
)

var (
	fontS = 12.0
	lineH = 14
)

type EquityBalancesResponse struct {
	EquityBalances   []*EquityBalance `json:"equity_balances"`
	EquityTotalValue *string          `json:"equity_total_value"`
	EquityTotalPL    *string          `json:"equity_total_pl"`
	EquityTotalPLP   *string          `json:"equity_total_plp"`
}

type EquityBalance struct {
	EquityBalanceID     *int64  `json:"equity_balance_id"`
	StockCode           *string `json:"stock_code"`
	StockName           *string `json:"stock_name"`
	MarketSection       *string `json:"market_section"`
	MarketSectionOLD    *string `json:"market_sector"` // FIXME: 後方互換性のため, to be removed
	TradeUnit           *string `json:"trade_unit"`
	AccountType         *string `json:"account_type"`
	BalanceQuantity     *string `json:"balance_quantity"`
	OrderingQuantity    *string `json:"ordering_quantity"`
	ShortableQuantity   *string `json:"shortable_quantity"`
	UnshortableQuantity *string `json:"unshortable_quantity"`
	BookUnitPrice       *string `json:"book_unit_price"`
	CurrentPrice        *string `json:"current_price"`
	CurrentPriceTime    *string `json:"current_price_time"`
	ReferencePrice      *string `json:"reference_price"`
	IsDelisted          bool    `json:"is_delisted"`
	IsLongable          bool    `json:"is_longable"`
	IsShortable         bool    `json:"is_shortable"`
}

func (w *Window) DrawInvAssetFrame() {
}

func (w *Window) DrawInvMarketFrame() {
}
func (w *Window) DrawInvPriceFrame() {
}
func (w *Window) DrawInvOrderFrame() {
}
func (w *Window) DrawInvBalDtlFrame() {
}
func (w *Window) DrawInvCBarFrame() {
}
func (w *Window) DrawInvCDtlFrame() {
}

func (w *Window) DrawInvEqbalFrame() {

	//表題エリア
	titleH := 20
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

func (w *Window) DrawAccCont(dx int, dy int) {
	w.DrawInvEqBalCont(GetTestDataEquityBalances(), dx, dy)

}

func (w *Window) DrawInvAssetCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invAssetWin.Win, dx, dy)
}

//現物残高
func (w *Window) DrawInvEqBalCont(eqbal *EquityBalancesResponse, dx int, dy int) {

	face := truetype.NewFace(App.Font, &truetype.Options{Size: fontS, DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawSimple(w.Win, App.invEqbalWin.Win, dx, dy)

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
func (w *Window) DrawInvMktCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invMktWin.Win, dx, dy)
}
func (w *Window) DrawInvPriceCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invPriceWin.Win, dx, dy)
}
func (w *Window) DrawInvOrderCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invOrderWin.Win, dx, dy)
}
func (w *Window) DrawInvBalDtlCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invBalDtlWin.Win, dx, dy)
}
func (w *Window) DrawInvCBarCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invCBarWin.Win, dx, dy)
}
func (w *Window) DrawInvCDtlCont(dx int, dy int) {
	tl.DrawSimple(w.Win, App.invCDtlWin.Win, dx, dy)
}
