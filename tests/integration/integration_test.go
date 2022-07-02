//go:build integration

package integration_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"

	"github.com/avoropaev/otus-go-banner-rotator/internal/apitest"
)

type IntegrationSuite struct {
	apitest.APISuite
}

func (s *IntegrationSuite) SetupSuite() {
	godotenv.Load("../../.env.test")

	apiURL := os.Getenv("BANNER_ROTATOR_GRPC_HOST") + ":" + os.Getenv("BANNER_ROTATOR_GRPC_PORT")

	s.Init(apiURL)
}

func (s *IntegrationSuite) TearDownSuite() {
	s.End()
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}
