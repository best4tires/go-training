package slices

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/best4tires/go-train/utils"
)

func TestClone(t *testing.T) {
	tests := []struct {
		in  []int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3},
		},
		{
			in:  nil,
			exp: nil,
		},
		{
			in:  []int{},
			exp: []int{},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Clone(test.in)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		in  []int
		q   int
		exp bool
	}{
		{
			in:  []int{1, 2, 3},
			q:   2,
			exp: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Contains(test.in, test.q)
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestFirstIndexOf(t *testing.T) {
	tests := []struct {
		in  []int
		q   int
		exp int
	}{
		{
			in:  []int{1, 2, 3},
			q:   2,
			exp: 1,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := FirstIndexOf(test.in, test.q)
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		in  []int
		idx int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			idx: 2,
			exp: []int{1, 2},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := DeleteAt(test.in, test.idx)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		in  []int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3},
		},
		{
			in:  []int{1, 2, 3, 1, 3, 2},
			exp: []int{1, 2, 3},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Deduplicate(test.in)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		cnt int
		in  []int
		exp []int
	}{
		{
			cnt: 2,
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3, 1, 2, 3},
		},
		{
			cnt: 5,
			in:  []int{12},
			exp: []int{12, 12, 12, 12, 12},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Repeat(test.cnt, test.in...)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		in     []int
		accept func(int) bool
		exp    []int
	}{
		{
			in:     []int{1, 2, 3},
			accept: func(n int) bool { return n%2 == 0 },
			exp:    []int{2},
		},
		{
			in:     []int{2, 4, 6, 8, 10},
			accept: func(n int) bool { return n > 6 },
			exp:    []int{8, 10},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Filter(test.in, test.accept)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestChunks(t *testing.T) {
	tests := []struct {
		in        []int
		chunkSize int
		exp       [][]int
	}{
		{
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			chunkSize: 3,
			exp: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			chunkSize: 3,
			exp: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
		{
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			chunkSize: 3,
			exp: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11},
			},
		},
		{
			in:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			chunkSize: 20,
			exp: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			in:        []int{},
			chunkSize: 5,
			exp:       [][]int{},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Chunks(test.in, test.chunkSize)
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want: %v, have %v", test.exp, res)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		in  []int
		old int
		new int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			old: 1,
			new: 2,
			exp: []int{2, 2, 3},
		},
		{
			in:  []int{0, 1, 0, 0, 1, 1, 1, 0, 1},
			old: 1,
			new: 2,
			exp: []int{0, 2, 0, 0, 2, 2, 2, 0, 2},
		},
		{
			in:  []int{11, 22, 66, 44, 55},
			old: 66,
			new: 33,
			exp: []int{11, 22, 33, 44, 55},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Replace(test.in, test.old, test.new)
			if utils.SlicesPtrEqual(test.in, res) {
				t.Fatalf("slices pointers are equal")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		in  []int
		exp int
	}{
		{
			in:  []int{1, 2, 3},
			exp: 6,
		},
		{
			in:  []int{11, 22, 33, 44},
			exp: 110,
		},
		{
			in:  []int{10, -5, 3},
			exp: 8,
		},
		{
			in:  []int{},
			exp: 0,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Sum(test.in)
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestAvg(t *testing.T) {
	tests := []struct {
		in  []int
		exp float64
	}{
		{
			in:  []int{1, 2, 3}, //6
			exp: 2.0,
		},
		{
			in:  []int{1, 3, 7, 24, 45}, //80
			exp: 16.0,
		},
		{
			in:  []int{5, 14, 22, 7, 11}, //59
			exp: 11.8,
		},
		{
			in:  []int{},
			exp: 0,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res := Avg(test.in)
			if utils.EpsEqual(1e-6, test.exp, res) {
				t.Fatalf("want %f, have %f", test.exp, res)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		in     []int
		exp    int
		experr bool
	}{
		{
			in:     []int{1, 2, 3},
			exp:    1,
			experr: false,
		},
		{
			in:     []int{167, 99, 234, 56, 421},
			exp:    56,
			experr: false,
		},
		{
			in:     []int{45, 122, 9, 333, 67},
			exp:    9,
			experr: false,
		},
		{
			in:     []int{},
			experr: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res, err := Min(test.in)
			if err != nil && !test.experr {
				t.Fatalf("your test had an unexpected error: %v", err)
			}
			if err == nil && test.experr {
				t.Fatalf("your test should have had an error but it didn't")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		in     []int
		exp    int
		experr bool
	}{
		{
			in:     []int{1, 2, 3},
			exp:    3,
			experr: false,
		},
		{
			in:     []int{167, 99, 234, 56, 421},
			exp:    421,
			experr: false,
		},
		{
			in:     []int{45, 122, 9, 333, 67},
			exp:    333,
			experr: false,
		},
		{
			in:     []int{},
			experr: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%02d", i), func(t *testing.T) {
			res, err := Max(test.in)
			if err != nil && !test.experr {
				t.Fatalf("your test had an unexpected error: %v", err)
			}
			if err == nil && test.experr {
				t.Fatalf("your test should have had an error but it didn't")
			}
			if !reflect.DeepEqual(test.exp, res) {
				t.Fatalf("want %v, have %v", test.exp, res)
			}
		})
	}
}
