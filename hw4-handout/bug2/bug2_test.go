package bug2

import (
	"sort"
	"testing"
)

func id(v int) int {
	return v
}

func double(v int) int {
	return 2 * v
}

// equal returns true if two slices are equal
func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func TestBug2(t *testing.T) {
	tests := []struct {
		n        int
		fun      func(int) int
		expected []int
	}{
		{5, id, []int{0, 1, 2, 3, 4}},
		{5, double, []int{0, 2, 4, 6, 8}},
	}

	for _, test := range tests {
		ch := make(chan int, 2)
		go bug2(test.n, test.fun, ch)
		// Store the output of the channel in a slice.
		var m []int
		for v := range ch {
			m = append(m, v)
		}
		// Sort the slice.
		sort.Ints(m)
		if !equal(test.expected, m) {
			t.Errorf("bug2 failed, expected %v, got %v", test.expected, m)
		}
	}
}
