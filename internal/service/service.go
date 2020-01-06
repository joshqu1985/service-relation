package service

import (
	"github.com/joshqu1985/service-relation/internal/dao/cache"
	"github.com/joshqu1985/service-relation/internal/dao/database"
)

type Service struct {
	Cache cache.Cache
	DB    database.DB
}

func New(c cache.Cache, db database.DB) *Service {
	return &Service{c, db}
}
