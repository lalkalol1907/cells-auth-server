package gRPCServer

import (
	"cells-auth-server/src/CustomErrors"
	"cells-auth-server/src/Repository"
	"cells-auth-server/src/gRPCServer/proto"
	"context"
	"encoding/json"
	"errors"
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
	if errors.Is(err, CustomErrors.NoSession) {
		e := "No session"
		return &proto.GetUserResponse{
			AuthOK:      false,
			Error:       &e,
			NeedRefresh: false,
		}, nil
	}
	if errors.Is(err, CustomErrors.NeedRefreshError) {
		return &proto.GetUserResponse{
			AuthOK:      false,
			NeedRefresh: true,
		}, nil
	}
	if err != nil {
		e := err.Error()
		return &proto.GetUserResponse{
			AuthOK:      false,
			Error:       &e,
			NeedRefresh: false,
		}, nil
	}

	userRes, err := json.Marshal(user)
	if err != nil {
		e := err.Error()
		return &proto.GetUserResponse{
			AuthOK:      false,
			Error:       &e,
			NeedRefresh: false,
		}, nil
	}

	u := string(userRes)

	return &proto.GetUserResponse{
		AuthOK:      true,
		NeedRefresh: false,
		User:        &u,
	}, nil
}
