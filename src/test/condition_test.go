package condition_test

import (
	"testing"
)

// func TestCondition(t *testing.T) {
// 	if v,err:someFun(); err == nil {
// 		t.Log("0")
// 	}else{
// 		t.Log("1")
// 	}
// }

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("未知")
		}
	}

}
