package grand

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{n: 100}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Digits(tt.args.n)
			fmt.Println("", res)
			assert.Equal(t, tt.want, len(res))
		})
	}
}

func TestLetters(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{n: 100}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Letters(tt.args.n)
			fmt.Println("", res)
			assert.Equal(t, tt.want, len(res))
		})
	}
}

func TestIntn(t *testing.T) {
	res := Intn(math.MaxInt64 / 10)
	fmt.Println("res", res)
}

func TestHashString(t *testing.T) {
	fmt.Println(HashString("a"))
}

func TestHashAndKeepRawPart(t *testing.T) {
	type args struct {
		s    string
		llen []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{"df", []int{10}},
			want: "df-32c220482c",
		},
		{
			args: args{"df222222222222222222", []int{10}},
			want: "df22222222-e6f75cbc50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, HashAndKeepRawPart(tt.args.s, tt.args.llen...))
		})
	}
}
