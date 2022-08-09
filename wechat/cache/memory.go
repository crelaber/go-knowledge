package cache

import (
	"sync"
	"time"
)

type Memory struct {
	sync.Mutex

	data map[string]*data
}

type data struct {
	Data   interface{}
	Expire time.Time
}

func NewMemory() *Memory {
	return &Memory{
		data: map[string]*data{},
	}
}

func (m Memory) Get(key string) interface{} {
	if ret, ok := m.data[key]; ok {
		if ret.Expire.Before(time.Now()) {
			m.deleteKey(key)
			return nil
		}
		return ret.Data
	}
	return nil
}

func (m Memory) Set(key string, val interface{}, timeout time.Duration) error {
	m.Lock()
	defer m.Unlock()
	m.data[key] = &data{
		Data:   val,
		Expire: time.Now().Add(timeout),
	}
	return nil
}

func (m Memory) IsExists(key string) bool {
	if ret, ok := m.data[key]; ok {
		if ret.Expire.Before(time.Now()) {
			m.deleteKey(key)
			return false
		}
		return true
	}
	return false
}

func (m Memory) Delete(key string) error {
	m.deleteKey(key)
	return nil
}

func (m Memory) deleteKey(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}
