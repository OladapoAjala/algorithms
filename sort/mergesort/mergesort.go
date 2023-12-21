package mergesort

import (
	"github.com/OladapoAjala/datastructures/sequences"
	"golang.org/x/exp/constraints"
)

/*
Paper ref.: https://www.researchgate.net/publication/312963714_Merge_sort_enhanced_in_place_sorting_algorithm
*/

func MergeSort[T constraints.Ordered](data sequences.Sequencer[T]) error {
	if data.GetSize() <= 1 {
		return nil
	}
	return split(data, 0, data.GetSize()-1)
}

func split[T constraints.Ordered](seq sequences.Sequencer[T], i, j int32) error {
	if j == i+1 || j == i {
		iVal, err := seq.GetData(i)
		if err != nil {
			return err
		}
		jVal, err := seq.GetData(j)
		if err != nil {
			return err
		}

		if iVal > jVal {
			return swap(seq, i, j)
		}
		return nil
	}

	mid := (i + j) / 2
	err := split(seq, i, mid)
	if err != nil {
		return err
	}
	err = split(seq, mid+1, j)
	if err != nil {
		return err
	}

	midData, err := seq.GetData(mid)
	if err != nil {
		return err
	}
	midIshData, err := seq.GetData(mid + 1)
	if err != nil {
		return err
	}

	if midIshData < midData {
		return sort(seq, i, j)
	}
	return nil
}

func sort[T constraints.Ordered](seq sequences.Sequencer[T], i, j int32) error {
	x := i
	a := x
	b := (i+j)/2 + 1
	y, ctr := b, b

	for x < b {
		ctrData, _ := seq.GetData(ctr)
		// if err != nil {
		// 	return err
		// }
		ctrIshData, _ := seq.GetData(ctr + 1)
		// if err != nil {
		// 	return err
		// }

		if ctr < j && ctrData > ctrIshData {
			swap(seq, ctr, ctr+1)
			ctr = ctr + 1
		}

		if ctr >= j || ctrData <= ctrIshData {
			ctr = b
		}

		if b > j && a > x && b == a+1 && a > y && ctr == b {
			b = a
			ctr = b
			a = y
		} else if b > j && a > x && ctr == b {
			b = y
			ctr = b
			a = x
		} else if b > j && ctr == b {
			break
		}

		if a == x && x == y && ctr == b {
			y = b
		} else if x == y {
			y = a
		}

		bData, err := seq.GetData(b)
		if err != nil {
			return err
		}
		aData, err := seq.GetData(a)
		if err != nil {
			return err
		}

		if a > y && b > a+1 && bData < aData && ctr == b {
			swap(seq, a, b)
			swap(seq, a, x)
			x = x + 1
			a = a + 1

			ctrData, err := seq.GetData(ctr)
			if err != nil {
				return err
			}
			ctrIshData, err := seq.GetData(ctr + 1)
			if err != nil {
				return err
			}

			if ctrData > ctrIshData {
				swap(seq, ctr, ctr+1)
				ctr = ctr + 1
			}
		} else if a == x && b == y && bData < aData {
			swap(seq, x, b)
			a = b
			b = b + 1
			x = x + 1
			if ctr == b-1 {
				ctr = ctr + 1
			}
		} else if a == x && b == y && bData >= aData {
			x = x + 1
			a = a + 1
		} else if b == a+1 && bData < aData {
			swap(seq, b, x)
			swap(seq, a, b)
			b = b + 1
			x = x + 1
			a = a + 1
			if ctr == b-1 {
				ctr = ctr + 1
			}
		} else if b == a+1 && bData >= aData {
			swap(seq, x, a)
			a = y
			x = x + 1
		} else if a == y && x < y && ctr == b+1 && bData < aData {
			swap(seq, x, b)
			b = b + 1
			x = x + 1
			if ctr == b-1 {
				ctr = ctr + 1
			}
		} else if b > a+1 && bData >= aData {
			swap(seq, x, a)
			x = x + 1
			a = a + 1
		}
	}
	return nil
}

func swap[T constraints.Ordered](seq sequences.Sequencer[T], i, j int32) error {
	iVal, err := seq.GetData(i)
	if err != nil {
		return err
	}
	jVal, err := seq.GetData(j)
	if err != nil {
		return err
	}

	err = seq.Insert(i, jVal)
	if err != nil {
		return err
	}
	err = seq.Insert(j, iVal)
	if err != nil {
		return err
	}
	return nil
}
