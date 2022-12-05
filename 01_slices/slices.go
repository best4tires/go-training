package slices

// Clone clones (returns a copy of) a slice
func Clone(ns []int) []int {
	//TODO: add implementation
	return nil
}

// Contains returns true if ns contains n, false otherwise
func Contains(ns []int, n int) bool {
	//TODO: add implementation
	return false
}

// FirstIndexOf return the first index of n if ns contains n, -1 otherwise
func FirstIndexOf(ns []int, n int) int {
	//TODO: add implementation
	return -1
}

// DeleteAt return a copy of ns with the element at idx deleted
func DeleteAt(ns []int, idx int) []int {
	//TODO: add implementation
	return nil
}

// Deduplicate returns a slice with all distinct values of ns (no duplicates)
func Deduplicate(ns []int) []int {
	//TODO: add implementation
	return nil
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

// Replace replaces a value of ns at idx with the new value q and returns a new slice
func Replace(ns []int, old int, new int) []int {
	// TODO: add implementation
	return nil
}

// Sum adds all values of ns together and returns the result. If the slice is empty it returns 0
func Sum(ns []int) int {
	// TODO: add implementation
	return 0
}

// Avg returns the avarage of all values of ns and returns the result. If the slice is empty it returns 0
func Avg(ns []int) float64 {
	// TODO: add implementation
	return 0
}

// Min returns the smallest value of ns
func Min(ns []int) (int, error) {
	// TODO: add implementation
	return 0, nil
}

// Max returns the biggest value of ns
func Max(ns []int) (int, error) {
	// TODO: add implementation
	return 0, nil
}
