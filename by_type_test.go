package gconv

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertTo(t *testing.T) {
	type args struct {
		v           interface{}
		toTypeValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "1", args: args{v: 1, toTypeValue: "222"}, want: "1"},
		{name: "2", args: args{v: "1", toTypeValue: 222}, want: 1},
		{name: "3", args: args{v: "1", toTypeValue: decimal.Zero}, want: decimal.NewFromInt(1)},
		{name: "4", args: args{v: "1", toTypeValue: nil}, want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ConvertTo(tt.args.v, tt.args.toTypeValue))
		})
	}
}
