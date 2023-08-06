package gos

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {

	got, err := Exec("sh", "-c", "tail -n 10 nohup.out > tmp && mv tmp nohup.out ")

	fmt.Println("got, err", got, err)
}
