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
	Fc        *Frame //チャート表示枠
	Fcr       *Frame //チャート実表示枠
	Fvr       *Frame //出来高表示枠
	cFntS     = 10
	prcRngTbl = []int{10, 20, 50, 100, 250, 500, 1000, 2500, 5000, 10000, 25000, 50000, 100000, 250000, 500000, 1000000, 2500000, 5000000, 10000000, 25000000, 50000000, 100000000}
	prcTbl    = []dayPrice{
		{"2018/10/4", 11205, 11415, 11050, 11200, 8303300, 11200},
		{"2018/10/3", 11250, 11255, 10965, 11055, 4302800, 11055},
		{"2018/10/2", 11455, 11465, 11170, 11190, 6350500, 11190},
		{"2018/10/1", 11380, 11500, 11270, 11435, 5436000, 11435},
		{"2018/9/28", 11250, 11500, 11150, 11470, 8708300, 11470},
		{"2018/9/27", 11115, 11185, 10960, 10960, 4412700, 10960},
		{"2018/9/26", 10955, 11220, 10935, 11065, 5386000, 11065},
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
		{"2018/8/24", 9867, 10100, 9770, 10100, 6056900, 10100},
		{"2018/8/23", 9820, 9969, 9809, 9854, 4318200, 9854},
		{"2018/8/22", 9530, 10015, 9521, 9950, 8470000, 9950},
		{"2018/8/21", 10000, 10015, 9770, 9821, 6552900, 9821},
		{"2018/8/20", 10080, 10115, 9970, 9984, 3725500, 9984},
		{"2018/8/17", 10080, 10150, 9985, 10020, 3585100, 10020},
		{"2018/8/16", 9944, 10125, 9888, 9985, 7882400, 9985},
		{"2018/8/15", 10375, 10385, 10150, 10170, 4067100, 10170},
		{"2018/8/14", 10180, 10470, 10175, 10445, 6262200, 10445},
		{"2018/8/13", 9940, 10155, 9921, 10070, 5479800, 10070},
		{"2018/8/10", 10445, 10560, 10050, 10120, 6477700, 10120},
		{"2018/8/9", 10305, 10540, 10275, 10490, 6571200, 10490},
		{"2018/8/8", 10190, 10600, 10160, 10530, 16308700, 10530},
		{"2018/8/7", 9930, 10105, 9842, 10050, 17130200, 10050},
		{"2018/8/6", 9275, 9445, 9252, 9433, 4565800, 9433},
		{"2018/8/3", 9300, 9329, 9193, 9232, 4581200, 9232},
		{"2018/8/2", 9336, 9353, 9191, 9235, 4564300, 9235},
		{"2018/8/1", 9390, 9400, 9292, 9367, 4624600, 9367},
		{"2018/7/31", 9200, 9360, 9135, 9260, 8316000, 9260},
		{"2018/7/30", 9235, 9312, 9183, 9276, 4693600, 9276},
		{"2018/7/27", 9250, 9385, 9221, 9385, 4397300, 9385},
		{"2018/7/26", 9451, 9489, 9252, 9254, 8373100, 9254},
		{"2018/7/25", 9597, 9645, 9555, 9570, 3757200, 9570},
		{"2018/7/24", 9637, 9739, 9547, 9598, 5657500, 9598},
		{"2018/7/23", 9722, 9779, 9537, 9571, 8477900, 9571},
		{"2018/7/20", 9700, 9888, 9601, 9857, 12345400, 9857},
		{"2018/7/19", 9600, 9829, 9518, 9758, 9793000, 9758},
		{"2018/7/18", 9730, 9800, 9601, 9650, 7221100, 9650},
		{"2018/7/17", 9720, 9909, 9602, 9604, 13011900, 9604},
		{"2018/7/13", 9500, 9750, 9433, 9722, 15999500, 9722},
		{"2018/7/12", 9020, 9413, 9015, 9376, 20449000, 9376},
		{"2018/7/11", 8725, 8831, 8698, 8812, 6911200, 8812},
		{"2018/7/10", 8728, 8820, 8716, 8758, 8446300, 8758},
		{"2018/7/9", 8351, 8585, 8350, 8578, 7068900, 8578},
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

	Fc = NewFrame(w.Fd.sx+35, w.Fd.sy+5, w.Fd.ex-35, w.Fd.ey-int(cFntS)-2) //チャートウインドウ表示領域
	Fcr = NewFrame(Fc.sx+1, Fc.sy+10, Fc.ex-1, Fc.sy+Fc.h*75/100)          //チャート 実表示領域(チャートウインドウの75%のサイズ)
	tl.DrawFrame(w.Win, cn.ColDarkGray, Fcr.sx, Fcr.sy, Fcr.ex, Fcr.ey, 0) //実チャート枠描画
	for y := Fcr.sy + 1; y < Fcr.ey-1; y++ {                               //背景色
		tl.DrawLine(w.Win, cn.ColGray11, Fcr.sx+1, y, Fcr.ex-1, y, 0)
	}

	Fvr = NewFrame(Fc.sx+1, Fc.sy+Fc.h*75/100+5, Fc.ex-1, Fc.ey)           //出来高 実表示領域(チャートウインドウの25% -5のサイズ)
	tl.DrawFrame(w.Win, cn.ColDarkGray, Fvr.sx, Fvr.sy, Fvr.ex, Fvr.ey, 0) //出来高枠描画描画
	for y := Fvr.sy + 1; y < Fvr.ey-1; y++ {                               //背景色
		tl.DrawLine(w.Win, cn.ColGray11, Fvr.sx+1, y, Fvr.ex-1, y, 0)
	}
}

//チャート
func (w *Window) DrawChartCont(dx int, dy int) {

	tl.DrawSimple(w.Win, App.chartWin.Win, dx, dy) //ベース描画

	if Fc.w < 150 || Fc.h < 80 { //ウインドウ幅が狭すぎる場合はこれ以上描画しない
		return
	}

	faceTitle := truetype.NewFace(App.Font, &truetype.Options{Size: float64(14.0), DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	tl.DrawText(w.Win, cn.ColWhite, faceTitle, dx+5, dy+14, fmt.Sprintf("9984 ソフトバンク"))

	face := truetype.NewFace(App.Font, &truetype.Options{Size: float64(cFntS), DPI: dpi, Hinting: font.HintingNone}) //font.HintingFull

	//TODO 25, 75日線描画

	w.drawChartBarCont(dx, dy, face) //チャート
	w.drawVolumeCont(dx, dy, face)   //出来高
}

func (w *Window) drawChartBarCont(dx int, dy int, face font.Face) { //チャート

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
		tl.DrawLine(w.Win, cn.ColGray66, dx+Fcr.sx, y, dx+Fcr.ex, y, chtLineGrid)
	}

	for i := 0; i < prcGridNum; i++ { //Y軸チャート指標値段描画
		y := dy + Fcr.ey - (Fcr.h/(prcGridNum-1))*i + 4
		vol := minCPrc + prcW*i
		if i == 0 { //位置調整
			y -= int(fontS) / 3
		} else if i == prcGridNum-1 {
			y += int(fontS) / 3
		}
		tl.DrawText(w.Win, cn.ColWhite, face, dx+2, y, fmt.Sprintf("%7d", vol))
		tl.DrawText(w.Win, cn.ColWhite, face, dx+Fcr.ex+3, y, fmt.Sprintf("%-7d", vol))
	}

	lastX := 10000
	for i, ch := range prcTbl { //X軸チャート指標描画（日時）＆指標線
		centerX := Fcr.ex - Fcr.w*(i+1)/(len(prcTbl)+1)
		if i != 0 && lastX < centerX+minCPrcInterval {
			continue
		}
		lastX = centerX
		tl.DrawText(w.Win, cn.ColWhite, face, dx+centerX-5, dy+Fc.ey+cFntS, fmt.Sprintf("%s", ch.date[5:])) //日時
		tl.DrawLine(w.Win, cn.ColGray66, dx+centerX, dy+Fcr.sy+1, dx+centerX, dy+Fcr.ey-1, 4)               //指標線
	}

	//ローソク足描画
	for i, ch := range prcTbl {

		base := float64(Fcr.h) / float64(prcRng)
		maxY := dy + Fcr.ey - int((ch.maxp-float64(minCPrc))*base)
		minY := dy + Fcr.ey - int((ch.minp-float64(minCPrc))*base)

		var startY, endY int
		if ch.startp > ch.endp {
			startY = dy + Fcr.ey - int((ch.startp-float64(minCPrc))*base)
			endY = dy + Fcr.ey - int((ch.endp-float64(minCPrc))*base)
		} else {
			startY = dy + Fcr.ey - int((ch.endp-float64(minCPrc))*base)
			endY = dy + Fcr.ey - int((ch.startp-float64(minCPrc))*base)
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

	mx, my := App.Ms.Px, App.Ms.Py
	if dx+Fcr.sx < mx && mx < dx+Fcr.ex && dy+Fcr.sy < my && my < dy+Fcr.ey { //カーソル位置線
		tl.DrawLine(w.Win, cn.ColWhite, mx, dy+Fcr.sy, mx, dy+Fcr.ey, 0) //X軸
		tl.DrawLine(w.Win, cn.ColWhite, dx+Fcr.sx, my, dx+Fcr.ex, my, 0) //Y軸
	}
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

	//	tl.DrawText(w.Win, cn.ColWhite, face, dx+Fvr.sx-int(fontS)-1, dy+Fvr.ey, fmt.Sprintf("%9d", 0))
	for i := 0; i < volGridNum; i++ { //出来高指標描画X
		y := dy + Fvr.ey - (Fvr.h/(volGridNum-1))*i + 4
		if i == 0 { //位置調整
			y -= int(fontS) / 3
		} else if i == volGridNum-1 {
			y += int(fontS) / 3
		}
		vol := maxDispVol / (volGridNum - 1) * (i) / int(math.Pow10(slen-2))
		tl.DrawText(w.Win, cn.ColWhite, face, dx+Fvr.sx-int(fontS)-2, y, fmt.Sprintf("%3d", vol))
		tl.DrawText(w.Win, cn.ColWhite, face, dx+Fvr.ex+3, y, fmt.Sprintf("%-3d", vol))
	}
	tl.DrawText(w.Win, cn.ColWhite, face, dx+1, dy+Fvr.ey+8, fmt.Sprintf("(x%d)", int(math.Pow10(slen-1))))

	lastX := 10000
	for i, _ := range prcTbl { //出来高指標線
		centerX := dx + Fvr.ex - Fvr.w*(i+1)/(len(prcTbl)+1)
		if i != 0 && lastX < centerX+minCPrcInterval {
			continue
		}
		lastX = centerX
		tl.DrawLine(w.Win, cn.ColGray66, dx+centerX, dy+Fvr.sy+1, dx+centerX, dy+Fvr.ey-1, 4) //指標線
	}

	//出来高線描画
	for i, ch := range prcTbl {
		volY := dy + Fvr.ey - (ch.quantity * (Fvr.h - 2) / maxVol)
		vWidth := getCandleWidth(len(prcTbl))

		centerX := dx + Fvr.ex - Fvr.w*(i+1)/(len(prcTbl)+1) //中央値X
		for x := centerX - vWidth; x <= centerX+vWidth; x++ {
			/*
				wdh := vWidth*2 + 1
						r, g, b, a := cn.ColDarkBlue.RGBA()
							a = uint32(float64(a) * float64(1-float64(x-(centerX-vWidth))/float64(wdh))) //グラデーション（ちょっとうまくいかない）
							tl.DrawLine(w.Win, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, x, dy+Fvr.ey-2, x, volY, 0)
							if a > 255 {
							a = 255
							}else if a <0 {
							 a = 0
							}
						}
			*/
			tl.DrawLine(w.Win, cn.ColDarkBlue, x, dy+Fvr.ey-2, x, volY, 0)
		}
	}

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
