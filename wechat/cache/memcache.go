package cache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type Memcache struct {
	conn *memcache.Client
}

func NewMemcache(server ...string) *Memcache {
	mc := memcache.New(server...)
	return &Memcache{mc}
}

func (m Memcache) Get(key string) interface{} {
	var err error
	var item *memcache.Item
	if item, err = m.conn.Get(key); err != nil {
		return nil
	}

	var result interface{}
	if err = json.Unmarshal(item.Value, &result); err != nil {
		return nil
	}
	return result
}

func (m Memcache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return err
	}
	item := &memcache.Item{
		Key:        key,
		Value:      data,
		Expiration: int32(timeout / time.Second),
	}
	return m.conn.Set(item)
}

func (m Memcache) IsExists(key string) bool {
	if _, err := m.conn.Get(key); err != nil {
		return false
	}
	return true
}

func (m Memcache) Delete(key string) (err error) {
	return m.conn.Delete(key)
}

//这句表示接口Cache只能有Memcache结构体实现
var _ Cache = (*Memcache)(nil)
