// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

package gconv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := Map(struct {
		A string
		B int
	}{A: "dfd", B: 2}, "B")
	fmt.Println("m", m)
	assert.Equal(t, m["B"], 2)
	assert.Nil(t, m["A"])
}

func TestMapExclude(t *testing.T) {
	m := MapExclude(struct {
		A string
		B int
	}{A: "dfd", B: 2}, "B")
	fmt.Println("m", m)
	assert.Equal(t, m["A"], "dfd")
	assert.Nil(t, m["B"])
}
