// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

// Package gconv implements powerful and convenient converting functionality for any types of variables.
//
// This package should keep much less dependencies with other packages.
package gconv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/errors"
)

func TestInt64E8WithPanic(t *testing.T) {
	testFunc := func() (ii int64, err error) {
		defer func() {
			if err2 := recover(); err2 != nil {
				fmt.Println("panic", err2)
				errNew := errors.Errorf("panic:%v", err2)
				err = errNew
				return
			}
		}()
		return Int64E8WithPanic(""), err
	}
	val, err := testFunc()
	if err != nil {
		fmt.Println("val.err", val, err.Error())
	}

	assert.Equal(t, int64(0), val)
	assert.NotNil(t, err)
}

func TestInt64E8(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "", args: args{any: "0.00000001"}, want: 1},
		{name: "", args: args{any: "0.00000002"}, want: 2},
		{name: "", args: args{any: "10000000000.00000001"}, want: 1000000000000000001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Int64E8(tt.args.any))
		})
	}
}

func TestInt64WithPanic(t *testing.T) {
	Int64WithPanic("1")
}
