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
		want func(int)
	}{
		{
			name: "simple tile packing",
			args: args{
				space: 3,
				tiles: []int{1, 2, 1},
			},
			want: func(perm int) {
				is.Equal(perm, 12)
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
