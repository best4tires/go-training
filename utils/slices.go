package utils

import (
	"reflect"
	"unsafe"
)

func SlicesPtrEqual[T any](ts1 []T, ts2 []T) bool {
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&ts1))
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&ts2))
	return h1.Data == h2.Data
}
