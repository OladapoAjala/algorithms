package mergesort

import (
	"testing"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/sequences/dynamicarray"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sequences/staticarray"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data sequences.Sequencer[string]
	}
	tests := []struct {
		name string
		args args
		want func(sequences.Sequencer[string], error)
	}{
		{
			name: "sort static array",
			args: args{
				data: dynamicarray.NewDynamicArray[string]("a", "f", "c", "m", "a", "e", "z"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				var i int32 = 0
				for i < sorted.GetSize()-1 {
					curVal, err := sorted.GetData(i)
					is.Nil(err)
					nxtVal, err := sorted.GetData(i + 1)
					is.Nil(err)
					is.LessOrEqual(curVal, nxtVal)
					i++
				}
			},
		},
		{
			name: "sort dynamic array",
			args: args{
				data: dynamicarray.NewDynamicArray[string]("a", "f", "c", "m", "a", "e", "z"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				var i int32 = 0
				for i < sorted.GetSize()-1 {
					curVal, err := sorted.GetData(i)
					is.Nil(err)
					nxtVal, err := sorted.GetData(i + 1)
					is.Nil(err)
					is.LessOrEqual(curVal, nxtVal)
					i++
				}
			},
		},
		{
			name: "sort linkedlist array",
			args: args{
				data: linkedlist.NewList[string]("a", "f", "c", "m", "a", "e", "z"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				var i int32 = 0
				for i < sorted.GetSize()-1 {
					curVal, err := sorted.GetData(i)
					is.Nil(err)
					nxtVal, err := sorted.GetData(i + 1)
					is.Nil(err)
					is.LessOrEqual(curVal, nxtVal)
					i++
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MergeSort(tt.args.data)
			tt.want(tt.args.data, err)
		})
	}
}

func Test_swap(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data sequences.Sequencer[string]
		i    int32
		j    int32
	}
	tests := []struct {
		name string
		args args
		want func(sequences.Sequencer[string], error)
	}{
		{
			name: "simple swap",
			args: args{
				data: staticarray.NewStaticArray[string](7, "a", "f", "c", "m", "a", "e", "z"),
				i:    1,
				j:    2,
			},
			want: func(sw sequences.Sequencer[string], err error) {
				is.Nil(err)
				iVal, err := sw.GetData(1)
				is.Nil(err)
				jVal, err := sw.GetData(2)
				is.Nil(err)
				is.Equal("c", iVal)
				is.Equal("f", jVal)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := swap(tt.args.data, tt.args.i, tt.args.j)
			tt.want(tt.args.data, err)
		})
	}
}
