package array_test

import (
	"testing"
)

// func TestArray(t *testing.T) {
// 	a := [2]int{1, 2}
// 	t.Log(a[1])
// }

// func TestArrayTravel(t *testing.T) {
// 	a := [...]int{1, 2, 3, 4, 5}
// 	// for i := 0; i < len(a); i++ {
// 	// 	t.Log(a[i])
// 	// }
// 	// for idx, e := range a {
// 	// 	t.Log(idx, e)
// 	// }
// 	for _, e := range a {
// 		t.Log(e)
// 	}
// }

//截取数组

func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[1:2])
}
