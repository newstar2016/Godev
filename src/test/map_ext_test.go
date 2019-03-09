package map_ext_test

import (
	"testing"
)

func TestMapExt(t *testing.T) {
	//map中存方法，实现工厂模式
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
	//map实现set

}

func TestMapSet(t *testing.T) {
	//map实现set
	myset := map[int]string{}
	myset[1] = "one"
	if v, ok := myset[1]; ok {
		t.Log(v)
	} else {
		t.Log("值不存在")
	}
	//map的个数
	t.Log(len(myset))
	//添加新的key，获取key的值
	myset[2] = "two"
	t.Log(len(myset))
	t.Log(myset[2])
	//删除key，检查key是否还存在
	delete(myset, 2)
	if v, ok := myset[2]; ok {
		t.Log(v)
	} else {
		t.Log("键不存在")
	}
}
