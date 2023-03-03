package dec_util

import (
	"math"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/zhwei820/gconv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var AlmostEqualPrecision int32 = 7 // almost-equal-precision

func DecToE8Int(v decimal.Decimal) int64 {
	return v.Shift(8).IntPart()
}

// 金额 => 逗号分隔字符串
// https://wiki.mark1.dev/pages/viewpage.action?pageId=25956170
func DecimalToEnStr(v decimal.Decimal, usd ...bool) string {
	var formatAmount = func(intPart int64, decimals decimal.Decimal, place int) string {
		p := message.NewPrinter(language.English)
		sInt := strings.TrimLeft(p.Sprintf("%d", intPart), "0")
		f, _ := decimals.Float64()
		s := strings.TrimLeft(p.Sprintf("%."+gconv.String(place)+"f", f), "0")
		return sInt + s
	}
	if v.LessThan(decimal.NewFromFloat(0.01)) {
		v = v.Truncate(8)
	} else {
		v = v.Truncate(5)
	}
	if v.GreaterThanOrEqual(decimal.NewFromInt(1)) {
		if len(usd) > 0 && usd[0] {
			v = v.Truncate(2)
			return formatAmount(v.IntPart(), v.Sub(v.Floor()), 2)
		} else {
			v = v.Truncate(5)
			s := formatAmount(v.IntPart(), v.Sub(v.Floor()), 5)
			return strings.TrimSuffix(s, "000")
		}
	}
	return v.String()
}

func FromE8Int(v int64) decimal.Decimal {
	return decimal.New(v, -8)
}

func FromString(v string) decimal.Decimal {
	d, _ := decimal.NewFromString(v)
	return d
}

func AlmostEqual(v1, v2 decimal.Decimal, place ...int32) bool {
	if len(place) == 0 {
		place = []int32{AlmostEqualPrecision}
	}
	return (v1.Sub(v2)).Abs().LessThanOrEqual(decimal.New(1, -place[0]))
}

func PowDecimal(a, b decimal.Decimal) decimal.Decimal {
	aa, _ := a.Float64()
	bb, _ := b.Float64()
	return decimal.NewFromFloat(math.Pow(aa, bb))
}

func Between(i, a, b decimal.Decimal) bool {
	return i.GreaterThanOrEqual(a) && i.LessThanOrEqual(b)
}
