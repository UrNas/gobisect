package gobisect

import "sort"

// InsertInt return slice of int after insert value in specific index
func InsertInt(slice []int, idx, value int) []int {
	newSliceLen := len(slice) + 1
	if idx >= newSliceLen {
		panic("slice bounds out of range")
	}
	newSlice := make([]int, newSliceLen)
	copy(newSlice[:idx], slice[:idx])
	copy(newSlice[idx+1:], slice[idx:])
	newSlice[idx] = value
	return newSlice
}

// InsertString return slice of strings after insert value in specific index
func InsertString(s []string, idx int, value string) []string {
	newSliceLen := len(s) + 1
	if idx >= newSliceLen {
		panic("slice bounds out of range")
	}
	newSlice := make([]string, newSliceLen)
	copy(newSlice[:idx], s[:idx])
	copy(newSlice[idx+1:], s[idx:])
	newSlice[idx] = value
	return newSlice
}

// BisecRightInt Return the index where to insert item value in slice,
func BisecRightInt(slice []int, value int) int {
	sort.Ints(slice)
	hi := len(slice)
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if value < slice[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// BisecLeftInt Return the index where to insert item value in slice.
func BisecLeftInt(slice []int, value int) int {
	sort.Ints(slice)
	hi := len(slice)
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if slice[mid] < value {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// InsortRightInt return slice of int's after Insert item value in slice, and keep it sorted
func InsortRightInt(slice []int, value int) []int {
	idx := BisecRightInt(slice, value)
	return InsertInt(slice, idx, value)
}

// InsortLeftInt return slice of int's after Insert item value in slice, and keep it sorted
func InsortLeftInt(slice []int, value int) []int {
	idx := BisecLeftInt(slice, value)
	return InsertInt(slice, idx, value)
}
