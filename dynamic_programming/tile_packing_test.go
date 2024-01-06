package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_packTile(t *testing.T) {
	is := assert.New(t)

	type args struct {
		space int
		tiles []int
	}
	tests := []struct {
		name string
		args args
		want func(float64)
	}{
		{
			name: "simple tile packing",
			args: args{
				space: 3,
				tiles: []int{1, 2, 1},
			},
			want: func(perm float64) {
				is.EqualValues(perm, 12)
			},
		},
		{
			name: "mid complex tile packing",
			args: args{
				space: 5,
				tiles: []int{1, 2, 1},
			},
			want: func(perm float64) {
				is.EqualValues(perm, 70)
			},
		},
		{
			name: "complex tile packing",
			args: args{
				space: 50,
				tiles: []int{1, 2, 1, 4},
			},
			want: func(perm float64) {
				is.EqualValues(perm, 3.536287775945517e+19)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := packTile(tt.args.space, tt.args.tiles)
			tt.want(got)
		})
	}
}
