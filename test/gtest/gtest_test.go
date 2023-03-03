// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

// Package gtest provides convenient test utilities for unit testing.
package gtest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gconv"
)

func Test_SliceStr(t *testing.T) {
	res := gconv.SliceStr([]int64{1, 2, 3})
	fmt.Println(res)
	assert.Equal(t, res, []string{"1", "2", "3"})
}
