package internalgrpc

import (
	"context"
	"net"
	"strconv"

	grpcZerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
	"github.com/avoropaev/otus-go-banner-rotator/internal/server"
)

type serv struct {
	host string
	port int
	app  app.Application

	server *grpc.Server
}

var _ server.IServer = (*serv)(nil)

func NewServer(host string, port int, app app.Application) server.IServer {
	return &serv{host, port, app, nil}
}

func (s *serv) Start(_ context.Context) error {
	lsn, err := net.Listen("tcp", s.host+":"+strconv.Itoa(s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			logging.StreamServerInterceptor(grpcZerolog.InterceptorLogger(log.Logger)),
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			logging.UnaryServerInterceptor(grpcZerolog.InterceptorLogger(log.Logger)),
		)),
	)

	return grpcServer.Serve(lsn)
}

func (s *serv) Stop(_ context.Context) error {
	if s.server == nil {
		return nil
	}

	s.server.GracefulStop()

	return nil
}
