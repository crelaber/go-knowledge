package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	ctx  context.Context
	conn redis.UniversalClient
}

type RedisOpts struct {
	Host        string `json:"host" yaml:"host"`
	Password    string `json:"password" yaml:"password"`
	Database    int    `json:"database" yaml:"database"`
	MaxIdle     int    `json:"max_idle" yaml:"max_idle"`
	MaxActive   int    `json:"max_active" yaml:"max_active"`
	IdleTimeout int    `json:"idle_timeout" yaml:"idle_timeout"`
}

func NewRedis(ctx context.Context, opts *RedisOpts) *Redis {
	conn := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        []string{opts.Host},
		DB:           opts.Database,
		Password:     opts.Password,
		IdleTimeout:  time.Second * time.Duration(opts.IdleTimeout),
		MinIdleConns: opts.MaxIdle,
	})
	return &Redis{
		conn: conn,
		ctx:  ctx,
	}
}

func (r *Redis) SetConn(conn redis.UniversalClient) {
	r.conn = conn
}

func (r *Redis) SetRedisCtx(ctx context.Context) {
	r.ctx = ctx
}

func (r *Redis) Get(key string) interface{} {
	result, err := r.conn.Do(r.ctx, "GET", key).Result()
	if err != nil {
		return nil
	}
	return result
}

func (r *Redis) Set(key string, val interface{}, timeout time.Duration) error {
	return r.conn.SetEX(r.ctx, key, val, timeout).Err()
}

func (r *Redis) IsExists(key string) bool {
	result, _ := r.conn.Exists(r.ctx, key).Result()
	return result > 0
}

func (r *Redis) Delete(key string) error {
	return r.conn.Del(r.ctx, key).Err()
}
