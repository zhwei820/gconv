package gconv

import (
	"fmt"
	"testing"
)

func Test_MapM(t *testing.T) {
	m := map[string]interface{}{"aa": 1, "bb": 2}
	m1 := Map(m, "aa")
	fmt.Println("m1", Export(m1))
}
