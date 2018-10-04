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

	baseScn  *Window //ベース画面
	outScn   *Window //表示用
	mMenu    *Window //メインメニュー
	eqbalWin *Window //株式残高
	chartWin *Window //チャートウインドウ

	updTime time.Time //更新時間

	Ms Mouse //マウス
}{}

type Mouse struct {
	Px, Py int
}
