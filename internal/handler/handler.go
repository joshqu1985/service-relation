package handler

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/joshqu1985/protos/pkg/pb"
	"github.com/joshqu1985/service-relation/internal/cache"
	"github.com/joshqu1985/service-relation/internal/store"
)

type SvrHandler struct {
	Log   *zap.SugaredLogger
	Cache cache.Cache
	Store store.Store
}

func RegisterHandler(svr *grpc.Server, cache cache.Cache, store store.Store,
	logger *zap.SugaredLogger) {
	pb.RegisterRelationServer(svr, &SvrHandler{
		Log:   logger,
		Cache: cache,
		Store: store})
}

func (s *SvrHandler) ListFollowers(ctx context.Context, in *pb.RelationListArgs) (*pb.RelationListInfo, error) {
	return &pb.RelationListInfo{}, nil
}
