package slices

import (
	"errors"
)

// Clone clones (returns a copy of) a slice
func Clone(ns []int) []int {
	if ns == nil {
		return nil
	}
	newNs := []int{}
	for _, v := range ns {
		newNs = append(newNs, v)
	}
	return newNs
}

// Contains returns true if ns contains n, false otherwise
func Contains(ns []int, n int) bool {
	for _, v := range ns {
		if v == n {
			return true
		}
	}
	return false
}

// FirstIndexOf return the first index of n if ns contains n, -1 otherwise
func FirstIndexOf(ns []int, n int) int {
	for i, v := range ns {
		if v == n {
			return i
		}
	}
	return -1
}

// DeleteAt return a copy of ns with the element at idx deleted
func DeleteAt(ns []int, idx int) []int {
	newNS := []int{}

	for i, v := range ns {
		if i != idx {
			newNS = append(newNS, v)
		}
	}

	return newNS
}

// Deduplicate returns a slice with all distinct values of ns (no duplicates)
func Deduplicate(ns []int) []int {
	newNS := []int{}

	for _, v := range ns {
		if !Contains(newNS, v) {
			newNS = append(newNS, v)
		}
	}

	return newNS
}

// Repeat returns a slice, where ns is cnt times repeated.
// It uses a variadic parameter (...int),to allow for p.e. Repeat(10, 2) or Repeat(10, 2, 5, 9)
func Repeat(cnt int, ns ...int) []int {
	slc := Clone(ns)
	original := Clone(ns)

	for i := 1; i < cnt; i++ {
		slc = append(slc, original...)
	}

	return slc
}

// Filter returns a slice with only values from ns accepted by accept
func Filter(ns []int, accept func(int) bool) []int {
	newNS := []int{}
	for _, v := range ns {
		if accept(v) {
			newNS = append(newNS, v)
		}
	}

	return newNS
}

// Chunks returns a slice of slices made of chunks of ns, where the size of each chunk is equal to chunkSize.
// Except the last chunk may contain less elements
func Chunks(ns []int, chunkSize int) [][]int {
	chunks := [][]int{}
	numberOfChunks := len(ns) / chunkSize

	if len(ns)%chunkSize != 0 {
		numberOfChunks++
	}

	if ns == nil {
		return chunks
	}
	nsElement := 0
	for i := 0; i < numberOfChunks; i++ {
		newChunk := []int{}
		for j := 0; j < chunkSize; j++ {
			if nsElement+1 > len(ns) {
				break
			}
			newChunk = append(newChunk, ns[nsElement])
			nsElement++
		}
		chunks = append(chunks, newChunk)
	}

	return chunks
}

//Replace replaces a value of ns at idx with the new value q and returns a new slice
func Replace(ns []int, old int, new int) []int {
	if ns == nil {
		return nil
	}

	newNs := Clone(ns)
	for i, v := range newNs {
		if v == old {
			newNs[i] = new
		}
	}
	return newNs
}

//Sum adds all values of ns together and returns the result. If the slice is empty it returns 0
func Sum(ns []int) int {
	if ns == nil {
		return 0
	}

	sum := 0
	for _, v := range ns {
		sum = sum + v
	}
	return sum
}

//Avg returns the avarage of all values of ns and returns the result. If the slice is empty it returns 0
func Avg(ns []int) float64 {
	if ns == nil {
		return 0
	}

	sum := Sum(ns)
	res := float64(sum) / float64(len(ns))
	return res
}

//Min returns the smallest value of ns
func Min(ns []int) (int, error) {
	if ns == nil {
		return 0, errors.New("slice had no values")
	}

	var smallNum int
	for i, v := range ns {
		if i == 0 {
			smallNum = v
		}
		if v < smallNum {
			smallNum = v
		}
	}
	return smallNum, nil
}

//Max returns the biggest value of ns
func Max(ns []int) (int, error) {
	if ns == nil {
		return 0, errors.New("slice had no values")
	}

	var bigNum int
	for _, v := range ns {
		if v > bigNum {
			bigNum = v
		}
	}
	return bigNum, nil
}
