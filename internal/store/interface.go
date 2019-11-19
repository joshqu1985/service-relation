package store

//go:generate mockgen -destination=../mock/store_mock.go -package=mock github.com/joshqu1985/service-relation/internal/store Store

import (
	"context"

	"github.com/joshqu1985/service-relation/internal/model"
)

type Store interface {
	ListFollowers(ctx context.Context, uid string) ([]*model.Follower, error)
	ListFollowings(ctx context.Context, uid string) ([]*model.Following, error)
}
