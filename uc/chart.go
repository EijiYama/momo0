package uc

import (
	"fmt"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/momo0/test001/cn"
	"github.com/momo0/test001/tl"
	"golang.org/x/image/font"
)

const (
	dpi             = 72.0 //フォントgpi
	minCPrcInterval = 100  //最小チャート価格間隔
	prcGridNum      = 6    //チャート価格グリッド線数
	chtLineGrid     = 4    //チャート価格線グリッド幅
	volLineGridW    = 4    //出来高線グリッド幅
)

var (
	Fc        Frame //チャート表示枠
	Fcr       Frame //チャート実表示枠
	Fvr       Frame //出来高表示枠
	cFntS     = 8
	prcRngTbl = []int{10, 20, 50, 100, 250, 500, 1000, 2500, 5000, 10000, 25000, 50000, 100000, 250000, 500000, 1000000, 2500000, 5000000, 10000000, 25000000, 50000000, 100000000}
	prcTbl    = []dayPrice{
		{"2018/9/25", 11015, 11140, 10905, 10925, 4891000, 10925},
		{"2018/9/21", 11100, 11250, 10935, 11045, 8766600, 11045},
		{"2018/9/20", 11105, 11115, 10815, 10880, 4872500, 10880},
		{"2018/9/19", 10880, 11115, 10860, 10915, 6003300, 10915},
		{"2018/9/18", 10800, 10820, 10585, 10710, 6431500, 10710},
		{"2018/9/14", 11095, 11100, 10685, 10945, 9989300, 10945},
		{"2018/9/13", 10640, 11000, 10605, 10990, 8595500, 10990},
		{"2018/9/12", 10260, 10550, 10250, 10495, 8316300, 10495},
		{"2018/9/11", 9922, 10150, 9899, 10150, 5044300, 10150},
		{"2018/9/10", 9850, 10010, 9828, 9906, 4090000, 9906},
		{"2018/9/7", 9912, 10045, 9893, 9940, 5176800, 9940},
		{"2018/9/6", 9740, 10040, 9717, 10005, 5564200, 10005},
		{"2018/9/5", 10110, 10125, 9849, 9850, 6444500, 9850},
		{"2018/9/4", 10295, 10325, 10160, 10245, 2707600, 10245},
		{"2018/9/3", 10250, 10335, 10175, 10250, 2862600, 10250},
		{"2018/8/31", 10250, 10360, 10235, 10300, 4444900, 10300},
		{"2018/8/30", 10055, 10195, 10030, 10195, 5189500, 10195},
		{"2018/8/29", 10290, 10315, 10070, 10100, 4723700, 10100},
		{"2018/8/28", 10295, 10365, 10230, 10235, 4728400, 10235},
		{"2018/8/27", 10100, 10280, 10090, 10175, 4748600, 10175},
	}
	mapVolDivNum = map[int]int{
		1:  2,
		2:  3,
		3:  4,
		4:  5,
		5:  6,
		6:  4,
		7:  3,
		8:  5,
		9:  4,
		10: 5,
	}
)

type dayPrice struct {
	date     string
	startp   float64
	maxp     float64
	minp     float64
	endp     float64
	quantity int
	fendp    float64
}

func (w *Window) DrawChartFrame() {

	Fc = Frame{sx: w.Fd.sx + 35, sy: w.Fd.sy + 5, ex: w.Fd.ex - 35, ey: w.Fd.ey - int(cFntS) - 2} //チャート＋出来高 表示領域
	Fc.w, Fc.h = Fc.ex-Fc.sx, Fc.ey-Fc.sy
	//	tl.DrawFrame(w.Win, cn.ColDarkGray, Fc.sx, Fc.sy, Fc.ex, Fc.ey, 0) //実チャート枠描画

	Fcr = Frame{sx: Fc.sx + 1, sy: Fc.sy + 10, ex: Fc.ex - 1, ey: Fc.sy + Fc.h*75/100} //チャート 実表示領域
	Fcr.w, Fcr.h = Fcr.ex-Fcr.sx, Fcr.ey-Fcr.sy
	tl.DrawFrame(w.Win, cn.ColDarkGray, Fcr.sx, Fcr.sy, Fcr.ex, Fcr.ey, 0) //実チャート枠描画
	for y := Fcr.sy + 1; y < Fcr.ey-1; y++ {                               //背景色
		tl.DrawLine(w.Win, cn.ColHHDarkGray, Fcr.sx+1, y, Fcr.ex-1, y, 0)
	}

	Fvr = Frame{sx: Fc.sx + 1, sy: Fc.sy + Fc.h*75/100 + 5, ex: Fc.ex - 1, ey: Fc.ey} //出来高 実表示領域
	Fvr.w, Fvr.h = Fvr.ex-Fvr.sx, Fvr.ey-Fvr.sy
	tl.DrawFrame(w.Win, cn.ColDarkGray, Fvr.sx, Fvr.sy, Fvr.ex, Fvr.ey, 0) //出来高枠描画描画
	for y := Fvr.sy + 1; y < Fvr.ey-1; y++ {                               //背景色
		tl.DrawLine(w.Win, cn.ColHHDarkGray, Fvr.sx+1, y, Fvr.ex-1, y, 0)
	}
}

//チャート
func (w *Window) DrawChartCont(dx int, dy int) {

	tl.DrawSimple(w.Win, App.chartWin.Win, dx, dy) //ベース描画

	if Fc.w < 150 || Fc.h < 80 { //ウインドウ幅が狭すぎる場合はこれ以上描画しない
		return
	}

	face := truetype.NewFace(App.Font, &truetype.Options{Size: float64(cFntS), DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawText(w.Win, cn.ColWhite, face, dx+5, dy+10, fmt.Sprintf("9984 ソフトバンク"))

	maxPrc, minPrc := 0.0, 0.0 //高値、安値
	for i, ch := range prcTbl {
		if i == 0 {
			maxPrc, minPrc = ch.maxp, ch.minp
		} else {
			if maxPrc < ch.maxp {
				maxPrc = ch.maxp
			}
			if minPrc > ch.minp {
				minPrc = ch.minp
			}
		}
	}
	diffPrc := maxPrc - minPrc //MAX高値
	if maxPrc == 0.0 || minPrc == 0.0 || diffPrc < 0.0 {
		return
	}

	prcRng := -1
	for i := 0; i < len(prcRngTbl); i++ { //見出し値段計算&描画	見出し値段を5分割作成
		if int(diffPrc) <= prcRngTbl[i] {
			prcRng = prcRngTbl[i]
			break
		}
	}
	if prcRng == -1 {
		return
	}
	prcW := prcRng / (prcGridNum - 1)
	minCPrc := int(minPrc) / prcW * prcW //prcWの最小値設定

	for i := 1; i < (prcGridNum - 1); i++ { //チャートグリッド線描画X
		y := dy + Fcr.ey - (Fcr.h/(prcGridNum-1))*i
		tl.DrawLine(w.Win, cn.ColDarkGray, dx+Fcr.sx, y, dx+Fcr.ex, y, chtLineGrid)
	}

	for i := 0; i < prcGridNum; i++ { //チャート指標値段描画X
		tl.DrawText(w.Win, cn.ColWhite, face, dx+1, dy+Fcr.ey-(Fcr.h/(prcGridNum-1))*i+4, fmt.Sprintf("%8d", minCPrc+prcW*i))
		tl.DrawText(w.Win, cn.ColWhite, face, dx+Fcr.ex+3, dy+Fcr.ey-(Fcr.h/(prcGridNum-1))*i+4, fmt.Sprintf("%-8d", minCPrc+prcW*i))
	}

	lastX := 10000
	for i, ch := range prcTbl { //チャート指標描画Y（日時）
		centerX := Fcr.ex - Fcr.w*(i+1)/(len(prcTbl)+1) - 5
		if i != 0 && i != len(prcTbl)-1 && lastX < centerX+minCPrcInterval {
			continue
		}
		lastX = centerX
		tl.DrawText(w.Win, cn.ColWhite, face, dx+centerX, dy+Fc.ey+cFntS, fmt.Sprintf("%s", ch.date[5:]))
	}

	//TODO 25, 75日線描画

	//ローソク足描画
	for i, ch := range prcTbl {

		maxY := dy + Fcr.ey - int((ch.maxp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))
		minY := dy + Fcr.ey - int((ch.minp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))

		var startY, endY int
		if ch.startp > ch.endp {
			startY = dy + Fcr.ey - int((ch.startp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))
			endY = dy + Fcr.ey - int((ch.endp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))
		} else {
			startY = dy + Fcr.ey - int((ch.endp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))
			endY = dy + Fcr.ey - int((ch.startp-float64(minCPrc))*float64(Fcr.h)/float64(prcRng))
		}

		gain := ch.startp < ch.endp
		col := cn.ColLightBlue
		if gain {
			col = cn.ColLightRed
		}

		cWidth := getCandleWidth(len(prcTbl))
		for y := maxY; y <= minY; y++ { //ローソク足描画
			centerX := dx + Fcr.ex - Fcr.w*(i+1)/(len(prcTbl)+1) //ローソクなので中央値X
			w.Win.SetRGBA(centerX, y, col)
			if startY <= y && y <= endY {
				for x := centerX - cWidth; x <= centerX+cWidth; x++ {
					if x == centerX {
						continue
					}
					w.Win.SetRGBA(x, y, col)
				}
			}
		}
	}

	w.drawVolumeCont(dx, dy, face) //出来高

}

// 出来高
func (w *Window) drawVolumeCont(dx int, dy int, face font.Face) {

	maxVol := 0
	for _, ch := range prcTbl {
		if ch.quantity > maxVol {
			maxVol = ch.quantity
		}
	}
	if maxVol == 0 {
		return
	}
	slen := len(fmt.Sprintf("%d", maxVol))
	topNum := maxVol/int(math.Pow10(slen-1)) + 1
	volGridNum := mapVolDivNum[topNum]
	maxDispVol := topNum * int(math.Pow10(slen-1)) //最大表示出来高 (2,234,567 の場合、3,000,000)
	//	fmt.Println("maxVol, maxDispVol, slen, int(math.Pow10(slen-1)) :", maxVol, maxDispVol, slen, int(math.Pow10(slen-1)))

	for i := 1; i < volGridNum-1; i++ { //チャートグリッド線描画X
		y := dy + Fvr.ey - (Fvr.h/(volGridNum-1))*i
		tl.DrawLine(w.Win, cn.ColDarkGray, dx+Fvr.sx, y, dx+Fvr.ex, y, volLineGridW)
	}

	tl.DrawText(w.Win, cn.ColWhite, face, dx+1, dy+Fvr.ey+4, fmt.Sprintf("%9d", 0))
	for i := 1; i < volGridNum; i++ { //チャート指標値段描画X
		tl.DrawText(w.Win, cn.ColWhite, face, dx+1, dy+Fvr.ey-(Fvr.h/(volGridNum-1))*i+4, fmt.Sprintf("%3d", maxDispVol/(volGridNum-1)*(i)/int(math.Pow10(slen-2))))
		tl.DrawText(w.Win, cn.ColWhite, face, dx+Fvr.ex+3, dy+Fvr.ey-(Fvr.h/(prcGridNum-1))*i+4, fmt.Sprintf("%-3d", maxDispVol/(volGridNum-1)*(i)/int(math.Pow10(slen-2))))
	}
	tl.DrawText(w.Win, cn.ColWhite, face, dx+1, dy+Fvr.ey+8, fmt.Sprintf("(x%d)", int(math.Pow10(slen-1))))

}

func getCandleWidth(pt int) int { //ローソク足の左右に加算する幅を計算

	w := Fcr.w/pt/2 - 2
	if w > 10 {
		w = 10
	} else if w < 1 {
		w = 1
	}

	return w
}
