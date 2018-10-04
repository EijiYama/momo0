package tl

import (
	"strings"
	"time"
)

var (
	mapText = map[string]string{
		//入出金サブタイプ
		CONV_CashTransferSubType + "00": "",
		CONV_CashTransferSubType + "02": "（発）委託保証金",
		CONV_CashTransferSubType + "03": "端株整理",
		CONV_CashTransferSubType + "06": "名義書換料",
		CONV_CashTransferSubType + "07": "（保）口座管理料",
		CONV_CashTransferSubType + "08": "（外国）口座管理料 ",
		CONV_CashTransferSubType + "09": "（金）口座管理料 ",
		CONV_CashTransferSubType + "11": "（株式先物）委託保証金 ",
		CONV_CashTransferSubType + "14": "（オプション）委託保証金 ",
		CONV_CashTransferSubType + "24": "銀行振込手数料 ",
		CONV_CashTransferSubType + "93": "その他預り金 ",
		CONV_CashTransferSubType + "99": "その他",
		//取引区分
		CONV_TradeType + "equity":       "現物",
		CONV_TradeType + "margin_open":  "信用新規",
		CONV_TradeType + "margin_close": "信用返済",
		CONV_TradeType + "margin_swap":  "現引現渡",
		// 売買区分
		CONV_OrderSide + "buy":  "買付",
		CONV_OrderSide + "sell": "売付",
		CONV_OrderSide + "in":   "入",
		CONV_OrderSide + "out":  "出",
		//注文処理状況
		CONV_OrderStatus + "new":            "注文中",
		CONV_OrderStatus + "acknowledged":   "板乗り済",
		CONV_OrderStatus + "partial_filled": "一部約定",
		CONV_OrderStatus + "filled":         "全部約定",
		CONV_OrderStatus + "cancelled":      "取消済",
		CONV_OrderStatus + "rejected":       "受付エラー等",
		CONV_OrderStatus + "expired":        "失効",
		//口座区分
		CONV_AccountType + "general":   "一般",
		CONV_AccountType + "specific":  "特定",
		CONV_AccountType + "exemptive": "非課税",
		//執行条件
		CONV_ExecutionType + "none":     "-",
		CONV_ExecutionType + "on_open":  "寄付",
		CONV_ExecutionType + "on_close": "引け",
		CONV_ExecutionType + "funari":   "不成",
		//発注条件
		CONV_OrderConditionType + "none": "-",
		CONV_OrderConditionType + "stop": "逆指値注文",
		//ストップ条件
		CONV_StopCondition + "upper": "以上",
		CONV_StopCondition + "lower": "以下",
		//有効期限タイプ
		CONV_ExpirationType + "day": "当日中",
		CONV_ExpirationType + "gtd": "期限付",
		//注文タイプ
		CONV_OrderType + "market": "成行",
		CONV_OrderType + "limit":  "指値",
		//取引履歴区分
		/*		CONV_TradeHistoryType + balUC.TradeTypeEquity:            "現物",
				CONV_TradeHistoryType + balUC.TradeTypeMarginOpen:        "信用新規",
				CONV_TradeHistoryType + balUC.TradeTypeMarginClose:       "信用返済",
				CONV_TradeHistoryType + balUC.TradeTypeMarginSwap:        "現引現渡",
				CONV_TradeHistoryType + balUC.TradeTypeWithdraw:          "出金",
				CONV_TradeHistoryType + balUC.TradeTypeDeposit:           "入金",
				CONV_TradeHistoryType + balUC.TradeTypeShipment:          "出庫",
				CONV_TradeHistoryType + balUC.TradeTypeReceipt:           "入庫",
				CONV_TradeHistoryType + balUC.TradeTypeDividend:          "配当金",
				CONV_TradeHistoryType + balUC.TradeTypeCapitalGainTax:    "譲渡益税収税",
				CONV_TradeHistoryType + balUC.TradeTypeCapitalGainRefund: "譲渡益税還付",
				CONV_TradeHistoryType + balUC.TradeTypeOther:             "その他",
				//受渡状況
				CONV_SettlementStatus + balUC.SettlementStatusSettled:   "受残",
				CONV_SettlementStatus + balUC.SettlementStatusUnSettled: "約残",
		*/                           //アクションタイプ
		CONV_ActionType + "new":     "新規",
		CONV_ActionType + "replace": "訂正",
		CONV_ActionType + "cancel":  "取消",
		//レポートタイプ
		CONV_ReportType + "acknowledged": "受付",
		CONV_ReportType + "partial_fill": "一部約定",
		CONV_ReportType + "fill":         "約定",
		CONV_ReportType + "cancelled":    "取消完了",
		CONV_ReportType + "replaced":     "訂正完了",
		CONV_ReportType + "rejected":     "エラー",
		CONV_ReportType + "expired":      "失効",
		//弁済区分
		CONV_MarginTradeType + "system":  "制度",
		CONV_MarginTradeType + "general": "一般",

		//性別
		CONV_Gender + "1": "男",
		CONV_Gender + "2": "女",
		//株式投資経験
		CONV_StockExp + "0": "-", //信用口座は未開設時にこの値が入る想定
		CONV_StockExp + "1": "経験なし",
		CONV_StockExp + "2": "1年未満",
		CONV_StockExp + "3": "1～3年",
		CONV_StockExp + "4": "3年以上",
		//口座種別
		CONV_Withholding + "0": "-", //信用口座は未開設時にこの値が入る想定
		CONV_Withholding + "1": "特定源徴なし",
		CONV_Withholding + "2": "特定源徴あり",
		CONV_Withholding + "3": "一般",
		//居住地国
		CONV_Residence + "0": "-",
		CONV_Residence + "1": "日本",
		CONV_Residence + "9": "日本以外",
		//運用スタイル
		CONV_Style + "1": "短期的運用",
		CONV_Style + "2": "長期的運用",
		//リスク
		CONV_Risk + "1": "ハイリスク",
		CONV_Risk + "2": "ミドルリスク",
		//証券会社勤務
		CONV_StockWorker + "0": "なし",
		CONV_StockWorker + "1": "あり",
		//上場企業可否
		CONV_ListedCompany + "0": "非上場企業",
		CONV_ListedCompany + "1": "上場企業",
		//続柄
		CONV_Relationship + "1": "本人",
		CONV_Relationship + "2": "配偶者",
		CONV_Relationship + "3": "子",
		CONV_Relationship + "4": "父・母",
		CONV_Relationship + "5": "その他",
		//国籍
		CONV_Nationality + "1": "日本国籍",
		CONV_Nationality + "9": "日本国籍以外",
		//FATCA
		CONV_Fatca + "0": "該当しない",
		CONV_Fatca + "1": "該当する",
		//Peps
		CONV_Peps + "0": "該当しない",
		CONV_Peps + "1": "該当する",
		//状況（審査）
		CONV_ReviewStatus + "0": "ROMユーザー",
		CONV_ReviewStatus + "1": "審査OK",
		CONV_ReviewStatus + "2": "審査NG",
		CONV_ReviewStatus + "3": "審査中",
		//徴求書類
		CONV_ConfirmDocType + "10": "免許証",
		CONV_ConfirmDocType + "03": "保険証",
		CONV_ConfirmDocType + "17": "パスポート",
		CONV_ConfirmDocType + "29": "マイナンバーカード",
		CONV_ConfirmDocType + "22": "在留カード",
		//職業
		CONV_Occupation + "75": "会社員",
		CONV_Occupation + "73": "会社役員",
		CONV_Occupation + "82": "公務員",
		CONV_Occupation + "59": "自営業",
		CONV_Occupation + "51": "医師",
		CONV_Occupation + "54": "弁護士・税理士",
		CONV_Occupation + "69": "パート・アルバイト",
		CONV_Occupation + "91": "学生",
		CONV_Occupation + "98": "無職・専業主婦",
		CONV_Occupation + "99": "その他",
		//取引動機
		CONV_TradingMotive + "1": "広告",
		CONV_TradingMotive + "5": "DM",
		CONV_TradingMotive + "6": "イベント",
		CONV_TradingMotive + "7": "親戚・知人",
		CONV_TradingMotive + "9": "その他",
		//年収
		CONV_Income + "1": "300万円未満",
		CONV_Income + "2": "300～500万円未満",
		CONV_Income + "3": "500～1000万円未満",
		CONV_Income + "4": "1000万円以上",
		//金融資産
		CONV_Asset + "1": "30万円未満",
		CONV_Asset + "2": "30～300万円未満",
		CONV_Asset + "3": "300万円以上",
		//配当金受取方法
		CONV_DividendPayee + "1": "銀行口座等受取",
		CONV_DividendPayee + "2": "証券口座受取",
		//預金種類
		CONV_DepositType + "1": "普通",
		CONV_DepositType + "2": "当座預金",
		CONV_DepositType + "5": "普通",
		//都道府県
		CONV_Prefecture + "01": "北海道",
		CONV_Prefecture + "02": "青森県",
		CONV_Prefecture + "03": "岩手県",
		CONV_Prefecture + "04": "宮城県",
		CONV_Prefecture + "05": "秋田県",
		CONV_Prefecture + "06": "山形県",
		CONV_Prefecture + "07": "福島県",
		CONV_Prefecture + "08": "茨城県",
		CONV_Prefecture + "09": "栃木県",
		CONV_Prefecture + "10": "群馬県",
		CONV_Prefecture + "11": "埼玉県",
		CONV_Prefecture + "12": "千葉県",
		CONV_Prefecture + "13": "東京都",
		CONV_Prefecture + "14": "神奈川県",
		CONV_Prefecture + "15": "新潟県",
		CONV_Prefecture + "16": "富山県",
		CONV_Prefecture + "17": "石川県",
		CONV_Prefecture + "18": "福井県",
		CONV_Prefecture + "19": "山梨県",
		CONV_Prefecture + "20": "長野県",
		CONV_Prefecture + "21": "岐阜県",
		CONV_Prefecture + "22": "静岡県",
		CONV_Prefecture + "23": "愛知県",
		CONV_Prefecture + "24": "三重県",
		CONV_Prefecture + "25": "滋賀県",
		CONV_Prefecture + "26": "京都府",
		CONV_Prefecture + "27": "大阪府",
		CONV_Prefecture + "28": "兵庫県",
		CONV_Prefecture + "29": "奈良県",
		CONV_Prefecture + "30": "和歌山県",
		CONV_Prefecture + "31": "鳥取県",
		CONV_Prefecture + "32": "島根県",
		CONV_Prefecture + "33": "岡山県",
		CONV_Prefecture + "34": "広島県",
		CONV_Prefecture + "35": "山口県",
		CONV_Prefecture + "36": "徳島県",
		CONV_Prefecture + "37": "香川県",
		CONV_Prefecture + "38": "愛媛県",
		CONV_Prefecture + "39": "高知県",
		CONV_Prefecture + "40": "福岡県",
		CONV_Prefecture + "41": "佐賀県",
		CONV_Prefecture + "42": "長崎県",
		CONV_Prefecture + "43": "熊本県",
		CONV_Prefecture + "44": "大分県",
		CONV_Prefecture + "45": "宮崎県",
		CONV_Prefecture + "46": "鹿児島県",
		CONV_Prefecture + "47": "沖縄県",
		//約定市場
		CONV_ExecutedMarket + "XTKS": "東証",
		CONV_ExecutedMarket + "XTK1": "ダークプール",
		//決済状況
		CONV_MarginSettlementStatus + "0": "決済済",
		CONV_MarginSettlementStatus + "1": "-",
		CONV_MarginSettlementStatus + "2": "決済中",
		//役職
		CONV_Posion + "1": "役員",
		CONV_Posion + "7": "幹部職員",
		CONV_Posion + "K": "特定部署",
		CONV_Posion + "C": "一般職員",
	}
)

const (
	CONV_CashTransferSubType    = "CashTransferSubType="    //入出金サブタイプ
	CONV_TradeType              = "TradeType="              //取引区分
	CONV_OrderSide              = "OrderSide="              // 売買区分
	CONV_OrderStatus            = "OrderStatus="            //注文処理状況
	CONV_AccountType            = "AccountType="            //口座区分
	CONV_ExecutionType          = "ExecutionType="          //執行条件
	CONV_OrderConditionType     = "OrderConditionType="     //発注条件
	CONV_StopCondition          = "StopCondition="          //ストップ条件
	CONV_ExpirationType         = "ExpirationType="         //有効期限タイプ
	CONV_OrderType              = "OrderType="              //注文タイプ
	CONV_TradeHistoryType       = "TradeHistoryType="       //取引履歴区分
	CONV_SettlementStatus       = "SettlementStatus="       //受渡状況
	CONV_ActionType             = "ActionType="             //アクションタイプ
	CONV_ReportType             = "ReportType="             //レポートタイプ
	CONV_MarginTradeType        = "MarginTradeType="        //弁済区分
	CONV_Gender                 = "Gender="                 //性別
	CONV_StockExp               = "StockExp="               //株式投資経験
	CONV_Withholding            = "Withholding="            //口座種別
	CONV_Residence              = "Residence="              //居住地国
	CONV_Style                  = "Style="                  //運用スタイル
	CONV_Risk                   = "Risk="                   //リスク
	CONV_StockWorker            = "StockWorker="            //証券会社勤務
	CONV_ListedCompany          = "ListedCompany="          //上場企業可否
	CONV_Relationship           = "Relationship="           //続柄
	CONV_Nationality            = "Nationality="            //国籍
	CONV_Fatca                  = "Fatca="                  //FATCA
	CONV_Peps                   = "Peps="                   //Peps
	CONV_ReviewStatus           = "Status="                 //状況（審査）
	CONV_ConfirmDocType         = "ConfirmDocType="         //徴求書類
	CONV_Occupation             = "Occupation="             //職業
	CONV_TradingMotive          = "TradingMotive="          //取引動機
	CONV_Income                 = "Income="                 //年収
	CONV_Asset                  = "Asset="                  //金融資産
	CONV_DividendPayee          = "DividendPayee="          //配当金受取方法
	CONV_DepositType            = "DepositType="            //預金種類
	CONV_Prefecture             = "Prefecture="             //都道府県
	CONV_ExecutedMarket         = "ExecutedMarket="         //約定市場
	CONV_MarginSettlementStatus = "MarginSettlementStatus=" //決済状況
	CONV_Posion                 = "Posion="                 //役職
)

func ConvName(key string, value *string) *string {

	if value == nil {
		return nil
	}

	out := mapText[key+*value]
	if out == "" {
		out = *value
	}

	//	fmt.Printf("ConvName key :[%s], in :[%s], out :[%s]\n", key, *value, out)
	return &out
}

func AddSlashToStringDate(stringDate string, trimZero bool) string {

	if len(stringDate) != 8 {
		return stringDate
	}
	slashedDate := stringDate[0:4] + "/"
	if stringDate[4:5] == "0" && trimZero {
		slashedDate += stringDate[5:6] + "/"
	} else {
		slashedDate += stringDate[4:6] + "/"
	}

	if stringDate[6:7] == "0" && trimZero {
		slashedDate += stringDate[7:8]
	} else {
		slashedDate += stringDate[6:8]
	}

	return slashedDate
}

func AddHyphenToStringDate(stringDate string, trimZero bool) string {

	if len(stringDate) != 8 {
		return stringDate
	}
	hyphenDate := stringDate[0:4] + "-"
	if stringDate[4:5] == "0" && trimZero {
		hyphenDate += stringDate[5:6] + "-"
	} else {
		hyphenDate += stringDate[4:6] + "-"
	}

	if stringDate[6:7] == "0" && trimZero {
		hyphenDate += stringDate[7:8]
	} else {
		hyphenDate += stringDate[6:8]
	}

	return hyphenDate
}

func AddHyphenToPStringDate(stringDate *string, trimZero bool) *string {
	if stringDate == nil {
		return nil
	}

	addedStrDt := AddHyphenToStringDate(PS2S(stringDate), trimZero)

	return &addedStrDt
}

func TrimPStringNumber(in *string) *string {

	rst := ""
	if in != nil {
		rst = TrimStringNumber(*in)
	}

	return &rst
}

func TrimStringNumber(in string) string {

	dotPt := strings.IndexAny(in, ".")
	if dotPt == -1 {
		return in
	}

	dec := strings.TrimRight(in, "0")
	dec = strings.TrimRight(dec, ".")

	return dec
}

func ConvertOrderTime(in *string) *string {
	if in == nil {
		return S2PS("")
	}

	timeUTC, err := time.Parse("2006-01-02T15:04:05Z07:00", *in)
	if err != nil {
		return in
	}
	timeJST := timeUTC.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	out := timeJST.Format("2006-01-02 15:04:05")

	return &out
}

func TimeToString(t time.Time, trimZero bool) string {

	if trimZero {
		return t.Format("2006-1-2 15:04:05.999")
	} else {
		return t.Format("2006-01-02 15:04:05.999")
	}
}

func PTimeToPString(t *time.Time, trimZero bool) *string {

	if t == nil {
		return nil
	}

	time := TimeToString(*t, trimZero)

	return &time
}
