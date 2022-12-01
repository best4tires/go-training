package slices

// Clone clones (returns a copy of) a slice
func Clone(ns []int) []int {
	newNs := []int{}
	if ns == nil {
		newNs = nil
		return newNs
	}

	for i := 0; i < len(ns); i++ {
		newNs = append(newNs, ns[i])
	}
	return newNs
}

// Contains returns true if ns contains n, false otherwise
func Contains(ns []int, n int) bool {
	//TODO: add implementation
	for i := 0; i < len(ns); i++ {
		if ns[i] == n {
			return true
		}
	}
	return false
}

// FirstIndexOf return the first index of n if ns contains n, -1 otherwise
func FirstIndexOf(ns []int, n int) int {
	for i := 0; i < len(ns); i++ {
		if ns[i] == n {
			return ns[0]
		}
	}
	return -1
}

// DeleteAt return a copy of ns with the element at idx deleted
func DeleteAt(ns []int, idx int) []int {
	//TODO: add implementation
	return nil
}

// Deduplicate returns a slice with all distinct values of ns (no duplicates)
func Deduplicate(ns []int) []int {
	// valueSaver := []int{}

	// for i := 0; i < len(ns); i++ {

	// }

	return ns
}

// Repeat returns a slice, where ns is cnt times repeated.
// It uses a variadic parameter (...int),to allow for p.e. Repeat(10, 2) or Repeat(10, 2, 5, 9)
func Repeat(cnt int, ns ...int) []int {
	// TODO: add implementation
	return nil
}

// Filter returns a slice with only values from ns accepted by accept
func Filter(ns []int, accept func(int) bool) []int {
	// TODO: add implementation
	return nil
}

// Chunks returns a slice of slices made of chunks of ns, where the size of each chunk is equal to chunkSize.
// Except the last chunk may contain less elements
func Chunks(ns []int, chunkSize int) [][]int {
	// TODO: add implementation
	return nil
}
