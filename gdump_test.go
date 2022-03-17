// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"fmt"
	"os"
	"testing"

	"github.com/shopspring/decimal"
)

func TestExport(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]interface{}{decimal.NewFromInt(100)}},
			want: "\"100\"\n",
		},
		{
			name: "",
			args: args{[]interface{}{map[string]interface{}{"dd": 3, "44": "fff"}}},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(EXPORT_KEY, "1")
			fmt.Println("", Export(tt.args.i...))
		})
	}
}

func TestNoExport(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[]interface{}{decimal.NewFromInt(100)}},
			want: "100",
		},
		{
			name: "",
			args: args{[]interface{}{map[string]interface{}{"dd": 3, "44": "fff"}}},
			want: "100",
		},
		{
			name: "",
			args: args{[]interface{}{}},
			want: "100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(EXPORT_KEY, "")
			fmt.Println("", Export(tt.args.i...))
		})
	}
}
