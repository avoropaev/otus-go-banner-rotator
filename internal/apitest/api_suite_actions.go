package apitest

import (
	"context"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/avoropaev/otus-go-banner-rotator/internal/server/pb"
)

type APISuiteActions struct {
	suite.Suite
	conn   *grpc.ClientConn
	client pb.BannerRotatorServiceClient
	ctx    context.Context
}

func (s *APISuiteActions) Init(apiURL string) {
	s.ctx = context.Background()

	ctx, cancel := context.WithTimeout(s.ctx, time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		apiURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.Require().NoError(err)

	s.conn = conn
	s.client = pb.NewBannerRotatorServiceClient(conn)
}

func (s *APISuiteActions) End() {
	if err := s.conn.Close(); err != nil {
		s.Require().NoError(err)
	}
}
