package gRPC

import (
	"cells-auth-server/src/Config"
	googleRPC "google.golang.org/grpc"
	"net"
)

func InitServer() error {
	listener, err := net.Listen("tcp", ":"+Config.Cfg.GrpcServer.Port)
	if err != nil {
		return err
	}

	s := googleRPC.NewServer()
	RegisterAuthServer(s, &GrpcServer{})
	if err := s.Serve(listener); err != nil {
		return err
	}

	return nil
}
