package slice_test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s0 := []int{1, 2, 3, 4}
	t.Log(len(s0), cap(s0))
}
