package sort

import "github.com/OladapoAjala/datastructures/sequences"

type Sorter[T sequences.Sequencer[T]] interface {
	Sort(T) T
}
