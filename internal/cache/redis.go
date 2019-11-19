package cache

import (
	"github.com/joshqu1985/fireman/pkg/redis"
)

type kv struct {
	Pool *redis.Pool
}

func NewRedisRepo(pool *redis.Pool) Cache {
	return &kv{Pool: pool}
}
