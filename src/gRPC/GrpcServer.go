package gRPC

import (
	"cells-auth-server/src/Repository"
	"context"
	"github.com/google/uuid"
)

type GrpcServer struct {
	UnimplementedAuthServer
}

func (s *GrpcServer) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	userUuid, err := uuid.Parse(in.AuthToken)
	if err != nil {
		return nil, err
	}

	user, err := Repository.GetUserBySession(userUuid)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		Uuid:     user.Uuid.String(),
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Nickname: user.Nickname,
	}, nil
}
