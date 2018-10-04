package tl

import (
	"strconv"
	"time"
)

func S2PS(s string) *string {
	return &s
}

func B2PB(b bool) *bool {
	return &b
}

func PB2PS(b *bool) *string {
	if b == nil {
		return nil
	}
	if *b == true {
		return S2PS("true")
	}
	return S2PS("false")
}

func B2PS(b bool) *string {
	if b == true {
		return S2PS("true")
	}
	return S2PS("false")
}

func PB2S(b *bool) string {
	if b == nil {
		return ""
	}
	return strconv.FormatBool(*b)
}

func I2PI(i int) *int {
	return &i
}

func I2PI64(i int64) *int64 {
	return &i
}

func I2PS(i int) *string {
	res := strconv.Itoa(i)
	return &res
}

func I2PS64(i int64) *string {
	res := strconv.FormatInt(i, 10)
	return &res
}

func PI2S(i *int) string {
	res := strconv.Itoa(*i)
	return res
}

func PI2PS(pi *int) *string {
	if pi == nil {
		return nil
	}

	res := strconv.Itoa(*pi)
	return &res

}

func PUI2PS64(pi *uint64) *string {
	if pi == nil {
		return nil
	}

	res := strconv.FormatUint(*pi, 10)
	return &res

}

func UI2PUI(u uint64) *uint64 {
	return &u
}

func TodayTimeString() string {
	loc := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(loc)
	return now.Format("20060102")
}

func JstNow() time.Time {
	return time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
}

func DateFormat(t time.Time) string {
	return t.Format("20060102")
}

func RFC3339String(t time.Time) string {
	return t.Format(time.RFC3339)
}

func PS2S(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func PI2I(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func PI2I64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func B2S(b bool) string {
	if b == true {
		return "true"
	}
	return "false"
}
