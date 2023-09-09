package gconv

import (
	"fmt"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	a := struct {
		A string
		B int
	}{
		A: "a",
		B: 1,
	}
	WriteToFile(a, "aa.json")
	a.A = ""
	a.B = 0
	LoadJsonFromFile(&a, "aa.json")
	fmt.Println("a", Export(a))
}

func TestPlainWriteToFile(t *testing.T) {
	a := struct {
		A string
		B int
	}{
		A: "a",
		B: 1,
	}
	WritePlainToFile(a, "aa.json")
	a.A = ""
	a.B = 0
	LoadPlainJsonFromFile(&a, "aa.json")
	fmt.Println("a", Export(a))
}
