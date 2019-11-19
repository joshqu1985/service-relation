package store

import (
	"github.com/joshqu1985/fireman/pkg/mysql"
)

type db struct {
	Pool *mysql.Pool
}

func NewMysqlRepo(pool *mysql.Pool) Store {
	return &db{Pool: pool}
}
