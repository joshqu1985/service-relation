package main

import (
	"fmt"
	"net"

	"github.com/joshqu1985/fireman/pkg/configor"
	"github.com/joshqu1985/fireman/pkg/discover"
	"github.com/joshqu1985/fireman/pkg/log"
	"github.com/joshqu1985/fireman/pkg/store/mysql"
	"github.com/joshqu1985/fireman/pkg/store/redis"
	"github.com/joshqu1985/fireman/pkg/tracing"
	"github.com/joshqu1985/fireman/pkg/transport/rpc"

	"github.com/joshqu1985/service-relation/internal/dao/cache"
	"github.com/joshqu1985/service-relation/internal/dao/database"
	"github.com/joshqu1985/service-relation/internal/handler"
	"github.com/joshqu1985/service-relation/internal/service"
)

type Config struct {
	Name     string
	Port     int
	Discover discover.Config
	Log      log.Config
	Redis    redis.Config
	Mysql    []mysql.Config
}

var (
	Conf Config
)

func init() {
	if err := configor.Load("./configs/conf.toml", &Conf); err != nil {
		panic(err)
	}

	log.Init(Conf.Log)

	if _, err := tracing.Init(Conf.Name); err != nil {
		panic(err)
	}

	if err := discover.Init(Conf.Discover); err != nil {
		panic(err)
	}
}

func main() {
	gsvr := rpc.NewUnaryServer()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", Conf.Port))
	if err != nil {
		panic(err)
	}

	if err := discover.Register(Conf.Name, Conf.Port); err != nil {
		panic(err)
	}

	s := service.New(cache.NewRepository(redis.NewPool(Conf.Redis)),
		database.NewRepository(mysql.NewPool(Conf.Mysql)))

	handler.RegisterHandler(gsvr, s)

	if err := gsvr.Serve(listener); err != nil {
		panic(err)
	}
}
