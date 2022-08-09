package lru

import (
	"reflect"
	"testing"
)

type KString string

func (d KString) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", KString("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(KString)) != "1234" {
		t.Fatalf("cache hit key1 = 1234 failed")
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

//测试，当使用内存超过了设定值时，是否会触发“无用”节点的移除：
func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, KString(v1))
	lru.Add(k2, KString(v2))
	lru.Add(k3, KString(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("RemoveOldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback)
	lru.Add("key1", KString("123456"))
	lru.Add("k2", KString("k2"))
	lru.Add("k3", KString("k3"))
	lru.Add("k4", KString("k4"))

	expect := []string{"key1", "k2"}
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("call OnEvited failed, expect keys equals to %s", expect)
	}
}
