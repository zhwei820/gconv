package gos

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {

	got, err := Exec("ls", "-ls", ".")
	fmt.Println("got, err", got, err)
}
