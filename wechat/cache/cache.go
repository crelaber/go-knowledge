package cache

import "time"

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExists(key string) bool
	Delete(key string) error
}
