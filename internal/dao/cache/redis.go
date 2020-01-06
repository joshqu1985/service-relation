package cache

import (
	"github.com/joshqu1985/fireman/pkg/store/redis"
)

type kv struct {
	Pool *redis.Pool
}

func NewRepository(pool *redis.Pool) Cache {
	return &kv{Pool: pool}
}
