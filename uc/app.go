package uc

import (
	"time"

	"github.com/golang/freetype/truetype"
)

var App = struct {
	Font     *truetype.Font
	Pix      []uint8
	selMMenu int

	WinSX, WinSY int //ウインドウサイズ

	baseScn *Window //ベース画面
	outScn  *Window //表示用
	mMenu   *Window //メインメニュー

	accEqbalWin *Window //口座残高

	invAssetWin  *Window //投資情報 預り資産
	invEqbalWin  *Window //投資情報
	invMktWin    *Window //投資情報
	invPriceWin  *Window //投資情報
	invOrderWin  *Window //投資情報
	invBalDtlWin *Window //投資情報
	invCBarWin   *Window //投資情報
	invCDtlWin   *Window //投資情報

	chartWin *Window //チャートウインドウ

	updTime time.Time //更新時間

	Ms Mouse //マウス
}{}

type Mouse struct {
	Px, Py int
}
