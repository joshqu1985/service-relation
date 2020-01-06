package database

import (
	"github.com/joshqu1985/fireman/pkg/store/mysql"
)

type db struct {
	Pool *mysql.Pool
}

func NewRepository(pool *mysql.Pool) DB {
	return &db{Pool: pool}
}
