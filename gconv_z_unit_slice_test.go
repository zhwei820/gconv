// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

package gconv_test

import (
	"testing"

	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gconv/test/gtest"
)

func Test_Slice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := 123.456
		t.AssertEQ(gconv.Bytes("123"), []byte("123"))
		// t.AssertEQ(gconv.Bytes([]interface{}{1}), []byte{1})
		// t.AssertEQ(gconv.Bytes([]interface{}{300}), []byte("[300]"))
		t.AssertEQ(gconv.Strings(value), []string{"123.456"})
		t.AssertEQ(gconv.Ints(value), []int{123})
		t.AssertEQ(gconv.Floats(value), []float64{123.456})
		t.AssertEQ(gconv.Interfaces(value), []interface{}{123.456})
	})

}

func Test_Slice_Empty(t *testing.T) {
	// Int.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Ints(""), []int{})
		t.Assert(gconv.Ints(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Int32s(""), []int32{})
		t.Assert(gconv.Int32s(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Int64s(""), []int64{})
		t.Assert(gconv.Int64s(nil), nil)
	})
	// Uint.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Uints(""), []uint{})
		t.Assert(gconv.Uints(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Uint32s(""), []uint32{})
		t.Assert(gconv.Uint32s(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Uint64s(""), []uint64{})
		t.Assert(gconv.Uint64s(nil), nil)
	})
	// Float.
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Floats(""), []float64{})
		t.Assert(gconv.Floats(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Float32s(""), []float32{})
		t.Assert(gconv.Float32s(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gconv.Float64s(""), []float64{})
		t.Assert(gconv.Float64s(nil), nil)
	})
}

func Test_Slice_Interfaces(t *testing.T) {
	// map
	gtest.C(t, func(t *gtest.T) {
		array := gconv.Interfaces(map[string]interface{}{
			"id":   1,
			"name": "john",
		})
		t.Assert(len(array), 1)
		t.Assert(array[0].(map[string]interface{})["id"], 1)
		t.Assert(array[0].(map[string]interface{})["name"], "john")
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			Id   int `json:"id"`
			Name string
		}
		array := gconv.Interfaces(&A{
			Id:   1,
			Name: "john",
		})
		t.Assert(len(array), 1)
		t.Assert(array[0].(*A).Id, 1)
		t.Assert(array[0].(*A).Name, "john")
	})
}

func Test_Slice_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int    `json:"id"`
		name string `json:"name"`
	}
	gtest.C(t, func(t *gtest.T) {
		user := &User{1, "john"}
		array := gconv.Interfaces(user)
		t.Assert(len(array), 1)
		t.Assert(array[0].(*User).Id, 1)
		t.Assert(array[0].(*User).name, "john")
	})
}
