package gRPCServer

import (
	"cells-auth-server/src/Config"
	"cells-auth-server/src/gRPCServer/proto"
	"google.golang.org/grpc"
	"net"
)

func InitServer() error {
	listener, err := net.Listen("tcp", ":"+Config.Cfg.GrpcServer.Port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	proto.RegisterAuthServer(s, &GrpcServer{})
	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}
