// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/gconv.

package gbinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeDecodeToFloat64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "", args: args{[]byte("")}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LeDecodeToFloat64(tt.args.b))
		})
	}
}
