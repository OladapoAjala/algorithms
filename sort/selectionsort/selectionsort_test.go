package selectionsort

import (
	"testing"

	"github.com/OladapoAjala/datastructures/sequences"
	"github.com/OladapoAjala/datastructures/sequences/dynamicarray"
	"github.com/OladapoAjala/datastructures/sequences/linkedlist"
	"github.com/OladapoAjala/datastructures/sequences/staticarray"
	"github.com/stretchr/testify/assert"
)

func Test_getLargest(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data sequences.Sequencer[string]
		lim  int32
	}
	tests := []struct {
		name string
		args args
		want func(int32, error)
	}{
		// TODO: test for error case
		{
			name: "get largest from static array",
			args: args{
				data: staticarray.NewStaticArray[string](5, "F", "H", "Z", "I", "A"),
				lim:  4,
			},
			want: func(largest int32, err error) {
				is.Nil(err)
				is.EqualValues(largest, 2)
			},
		},
		{
			name: "get largest from dynamic array",
			args: args{
				data: dynamicarray.NewDynamicArray[string]("F", "H", "Z", "I", "A"),
				lim:  4,
			},
			want: func(largest int32, err error) {
				is.Nil(err)
				is.EqualValues(largest, 2)
			},
		},
		{
			name: "get largest from linkedlist",
			args: args{
				data: linkedlist.NewList[string]("F", "H", "Z", "I", "A"),
				lim:  4,
			},
			want: func(largest int32, err error) {
				is.Nil(err)
				is.EqualValues(largest, 2)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLargest(tt.args.data, tt.args.lim)
			tt.want(got, err)
		})
	}
}

func Test_selectionSort(t *testing.T) {
	is := assert.New(t)

	type args struct {
		data sequences.Sequencer[string]
		end  int32
	}
	tests := []struct {
		name string
		args args
		want func(sequences.Sequencer[string], error)
	}{
		{
			name: "sort static array",
			args: args{
				data: staticarray.NewStaticArray[string](5, "F", "H", "Z", "I", "A"),
				end:  4,
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, staticarray.NewStaticArray[string](5, "A", "F", "H", "I", "Z"))
			},
		},
		{
			name: "sort dynamic array",
			args: args{
				data: dynamicarray.NewDynamicArray[string]("F", "H", "Z", "I", "A"),
				end:  4,
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, dynamicarray.NewDynamicArray[string]("A", "F", "H", "I", "Z"))
			},
		},
		{
			name: "get largest from linkedlist",
			args: args{
				data: linkedlist.NewList[string]("F", "H", "Z", "I", "A"),
				end:  4,
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, linkedlist.NewList[string]("A", "F", "H", "I", "Z"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := selectionSort(tt.args.data, tt.args.end)
			tt.want(got, err)
		})
	}
}

func TestSelectionSort(t *testing.T) {
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
				data: staticarray.NewStaticArray[string](5, "F", "H", "Z", "I", "A"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, staticarray.NewStaticArray[string](5, "A", "F", "H", "I", "Z"))
			},
		},
		{
			name: "sort dynamic array",
			args: args{
				data: dynamicarray.NewDynamicArray[string]("F", "H", "Z", "I", "A"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, dynamicarray.NewDynamicArray[string]("A", "F", "H", "I", "Z"))
			},
		},
		{
			name: "get largest from linkedlist",
			args: args{
				data: linkedlist.NewList[string]("F", "H", "Z", "I", "A"),
			},
			want: func(sorted sequences.Sequencer[string], err error) {
				is.Nil(err)
				is.Equal(sorted, linkedlist.NewList[string]("A", "F", "H", "I", "Z"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectionSort(tt.args.data)
			tt.want(got, err)
		})
	}
}
