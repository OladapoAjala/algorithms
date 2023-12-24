package selectionsort

import (
	"github.com/OladapoAjala/datastructures/sequences"
	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](data sequences.Sequencer[T]) (sequences.Sequencer[T], error) {
	if data.GetSize() <= 1 {
		return data, nil
	}

	return selectionSort[T](data, data.GetSize()-1)
}

func selectionSort[T constraints.Ordered](data sequences.Sequencer[T], end int32) (sequences.Sequencer[T], error) {
	if end == 0 {
		return data, nil
	}

	l, err := getLargest[T](data, end)
	if err != nil {
		return nil, err
	}

	// Get value at new largest index
	m, err := data.GetData(l)
	if err != nil {
		return nil, err
	}
	// Get value at end of array
	n, err := data.GetData(end)
	if err != nil {
		return nil, err
	}

	// Swap largest value with the value at the end of array
	err = data.Insert(end, m)
	if err != nil {
		return nil, err
	}
	err = data.Insert(l, n)
	if err != nil {
		return nil, err
	}

	return selectionSort[T](data, end-1)
}

func getLargest[T constraints.Ordered](data sequences.Sequencer[T], lim int32) (int32, error) {
	if lim == 0 {
		return lim, nil
	}

	l, err := getLargest[T](data, lim-1)
	if err != nil {
		return -1, nil
	}

	// Get value at new largest index
	m, err := data.GetData(l)
	if err != nil {
		return -1, err
	}

	// Get value at end of sequence
	n, err := data.GetData(lim)
	if err != nil {
		return -1, err
	}

	if m > n {
		return l, nil
	}

	return lim, nil
}
