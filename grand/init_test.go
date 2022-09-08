package grand

import (
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
			// fmt.Println("", res)
			assert.LessOrEqual(t, tt.want, len(res))
		})
	}
}
