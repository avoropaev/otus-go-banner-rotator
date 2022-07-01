package cmd

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/avoropaev/otus-go-banner-rotator/cmd/config"
	"github.com/avoropaev/otus-go-banner-rotator/internal/app"
	internalGRPC "github.com/avoropaev/otus-go-banner-rotator/internal/server/grpc"
	psqlStorage "github.com/avoropaev/otus-go-banner-rotator/internal/storage"
)

func serveHTTPCommand(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve-http",
		Short: "serves http api",
		RunE:  serveHTTPCommandRunE(ctx),
	}

	command.Flags().StringVar(&cfgFile, "config", "", "Path to configuration file")

	err := command.MarkFlagRequired("config")
	if err != nil {
		return nil
	}

	return command
}

func serveHTTPCommandRunE(ctx context.Context) func(cmd *cobra.Command, args []string) (err error) {
	return func(cmd *cobra.Command, args []string) (err error) {
		configFile := cmd.Flag("config").Value.String()

		cfg, err := config.ParseBannerRotatorConfig(configFile)
		if err != nil {
			log.Error().Err(err).Msg("failed to parse config")

			return err
		}

		logLevel, err := zerolog.ParseLevel(cfg.Logger.Level)
		if err != nil {
			log.Error().Err(err).Msg("failed to install log level")

			return err
		}

		zerolog.SetGlobalLevel(logLevel)

		ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		defer cancel()

		conn, err := pgxpool.Connect(ctx, cfg.DB.URI)
		if err != nil {
			log.Error().Err(err).Msg("unable to connect to database")

			return err
		}
		go func() {
			conn.Close()
		}()

		store := psqlStorage.New(conn)

		application := app.New(store)
		grpcServer := internalGRPC.NewServer(cfg.GRPC.Host, cfg.GRPC.Port, application)

		go func() {
			<-ctx.Done()

			log.Info().Msg("stopping an grpc server...")

			ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()

			if err := grpcServer.Stop(ctx); err != nil {
				log.Error().Err(err).Msg("failed to stop grpc server")
			}
		}()

		log.Info().Msg("banner-rotator is running...")

		if err := grpcServer.Start(ctx); err != nil {
			cancel()

			if !errors.Is(err, http.ErrServerClosed) {
				log.Error().Err(err).Msg("failed to start grpc server")
			}
		}

		return nil
	}
}
