package dec_util

import (
	"math"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/zhwei820/gconv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func DecToE8Int(v decimal.Decimal) int64 {
	return v.Shift(8).IntPart()
}

// 金额 => 逗号分隔字符串
//
// 大于1:
// u: 2位
// 非u：2～5位，即至少2位，超过2位末尾=0隐藏，小于2位补0
// 小于1: 2～8位，即至少2位，即至少2位，超过2位末尾=0隐藏，小于2位补0

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

func AlmostEqual(v1, v2 decimal.Decimal, place ...float64) bool {
	if len(place) == 0 {
		place = []float64{7}
	}
	return (v1.Sub(v2)).Abs().LessThanOrEqual(decimal.NewFromFloat(1. / math.Pow(10, place[0])))
}

func PowDecimal(a, b decimal.Decimal) decimal.Decimal {
	aa, _ := a.Float64()
	bb, _ := b.Float64()
	return decimal.NewFromFloat(math.Pow(aa, bb))
}
