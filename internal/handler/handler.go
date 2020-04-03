package handler

import (
	"context"

	"google.golang.org/grpc"

	"github.com/joshqu1985/protos/pb"

	"github.com/joshqu1985/service-relation/internal/service"
)

func RegisterHandler(grpcSvr *grpc.Server, s *service.Service) {
	pb.RegisterBoardServer(grpcSvr, &Handler{s})
}

type Handler struct {
	Service *service.Service
}

func (h *Handler) ListFollowers(ctx context.Context, in *pb.RelationListArgs) (*pb.RelationListInfo, error) {
	return &pb.RelationListInfo{}, nil
}
