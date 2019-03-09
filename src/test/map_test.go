package test

import (
	"testing"
)

func TestMap(t *testing.T) {
	//如何声明map,string代表key的类型，int代表value的类型
	//m := map[string]int{"one": 1, "two": 2, "three": 3}
	//m2 := make(map[string]int, 10)
	//map不能用cap函数来统计
	//t.Log(cap(m2))
	//map在key不存在时默认赋予0值，所以要注意判断
	// m3 := map[int]int{}
	// t.Log(m3[3])
	// m3[3] = 0
	// if v, ok := m3[3]; ok {
	// 	t.Log("key是3的值存在是：", v)
	// } else {
	// 	t.Log("key是3的不存在")
	// }

	//当key是four的时候，打印时第4个键值对会在第一个循环打印，不是four,或者大于4个，或者小于4个都能正常
	//主要原因是遍历map时，元素的顺序是随机的
	//遍历map中的值
	m4 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for key, v := range m4 {
		t.Log(key, "=", v)
	}
}
