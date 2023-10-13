package gRPCServer

import (
	"cells-auth-server/src/Repository"
	"cells-auth-server/src/gRPCServer/proto"
	"context"
	"github.com/google/uuid"
)

type GrpcServer struct {
	proto.UnimplementedAuthServer
}

func (s *GrpcServer) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userUuid, err := uuid.Parse(in.AuthToken)
	if err != nil {
		return nil, err
	}

	user, err := Repository.GetUserBySession(userUuid)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		Uuid:     user.Uuid.String(),
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Nickname: user.Nickname,
	}, nil
}
