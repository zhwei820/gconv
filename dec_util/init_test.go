package dec_util

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/shopspring/decimal"
)

func TestDecToE8Int(t *testing.T) {
	type args struct {
		v decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{v: decimal.NewFromFloat(1.11)},
			want: int64(111000000),
		},
		{
			name: "",
			args: args{v: FromString("578372667.847732751905580779")},
			want: int64(57837266784773275),
		},
		{
			name: "",
			args: args{v: FromString("272667.8477327519")},
			want: int64(27266784773275),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DecToE8Int(tt.args.v))
		})
	}
}

func TestAlmostEqual(t *testing.T) {
	type args struct {
		v1    decimal.Decimal
		v2    decimal.Decimal
		place []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				v1: decimal.NewFromFloat(0.1),
				v2: decimal.NewFromFloat(0.1000001),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				v1: decimal.NewFromFloat(0.1),
				v2: decimal.NewFromFloat(0.1001),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				v1: decimal.NewFromFloat(5972.4904126),
				v2: decimal.NewFromFloat(5972.490412590000000000),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AlmostEqual(tt.args.v1, tt.args.v2, tt.args.place...))
		})
	}
}

func TestDecimalToEnStr(t *testing.T) {
	type args struct {
		v decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: decimal.NewFromInt(10000)}, want: "10,000.00"},
		{name: "2", args: args{v: decimal.NewFromFloat(10000.1)}, want: "10,000.10"},
		{name: "3", args: args{v: decimal.NewFromFloat(0.0000123)}, want: "0.0000123"},
		{name: "3", args: args{v: decimal.NewFromFloat(1.0000123)}, want: "1.00001"},
		{name: "3", args: args{v: decimal.NewFromFloat(1.0000)}, want: "1.00"},
		{name: "3", args: args{v: FromString("1234610169845720.625378816191095892")}, want: "1,234,610,169,845,720.62537"},
		{name: "3", args: args{v: FromString("9223372036854775807.625378816191095892")}, want: "9,223,372,036,854,775,807.62537"},  // max for format
		{name: "3", args: args{v: FromString("9223372036854775808.625378816191095892")}, want: "-9,223,372,036,854,775,808.62537"}, // overflow
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DecimalToEnStr(tt.args.v))
		})
	}
}

func TestDecimalToEnStrUSD(t *testing.T) {
	type args struct {
		v decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: decimal.NewFromInt(10000)}, want: "10,000.00"},
		{name: "2", args: args{v: decimal.NewFromFloat(10000.1)}, want: "10,000.10"},
		{name: "3", args: args{v: decimal.NewFromFloat(0.0000123)}, want: "0.0000123"},
		{name: "3", args: args{v: decimal.NewFromFloat(1.0000123)}, want: "1.00"},
		{name: "3", args: args{v: decimal.NewFromFloat(1.0000)}, want: "1.00"},
		{name: "3", args: args{v: decimal.NewFromFloat(92236247320.366446472747643836)}, want: "92,236,247,320.36"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DecimalToEnStr(tt.args.v, true))
		})
	}
}
