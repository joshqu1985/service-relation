package cache

//go:generate mockgen -destination=../mock/cache_mock.go -package=mock github.com/joshqu1985/service-relation/internal/cache Cache

import (
	"context"

	"github.com/joshqu1985/service-relation/internal/model"
)

type Cache interface {
	GetCounter(ctx context.Context, uid string) (model.Counter, error)
	SetCounter(ctx context.Context, c model.Counter) error
}
