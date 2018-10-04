package uc

import (
	"github.com/momo0/test001/tl"
)

const (
	AccountTypeGeneral   = "general"   // 一般
	AccountTypeSpecific  = "specific"  // 特定
	AccountTypeExemptive = "exemptive" // 非課税
)

func GetTestDataEquityBalances() *EquityBalancesResponse {
	return &EquityBalancesResponse{
		EquityTotalValue: tl.S2PS("1000"),
		EquityTotalPL:    tl.S2PS("32893"),
		EquityTotalPLP:   tl.S2PS("1.125"),
		EquityBalances: []*EquityBalance{
			&EquityBalance{
				EquityBalanceID:     tl.I2PI64(1111),              //注文ID
				StockCode:           tl.S2PS("8601"),              // 銘柄コード
				StockName:           tl.S2PS("大和証券Ｇ本社"),           // 銘柄名
				AccountType:         tl.S2PS(AccountTypeSpecific), // 税区分
				BalanceQuantity:     tl.S2PS("2000"),              // 残高株数
				OrderingQuantity:    tl.S2PS("0"),                 // 注文中株数
				ShortableQuantity:   tl.S2PS("2000"),              // 売付可能株数
				UnshortableQuantity: tl.S2PS("2000"),              // 売付不能株数
				BookUnitPrice:       tl.S2PS("735.5"),             // (optional) 概算簿価単価
				CurrentPrice:        tl.S2PS("1603.5"),            // (optional) 現在価格
				IsDelisted:          false,                        // 上場廃止フラグ
				IsLongable:          false,                        // 買付可能フラグ
				IsShortable:         false,                        // 売付可能フラグ
			},
			&EquityBalance{
				EquityBalanceID:     tl.I2PI64(2222),
				StockCode:           tl.S2PS("8604"),
				StockName:           tl.S2PS("野村証券"),
				AccountType:         tl.S2PS(AccountTypeGeneral),
				BalanceQuantity:     tl.S2PS("2000"),
				OrderingQuantity:    tl.S2PS("0"),
				ShortableQuantity:   tl.S2PS("2000"),
				UnshortableQuantity: tl.S2PS("2000"),
				BookUnitPrice:       tl.S2PS("735.5"),
				IsDelisted:          false,
				IsLongable:          true,
				IsShortable:         true,
			},
		},
	}
}
