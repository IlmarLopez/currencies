package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/IlmarLopez/currency/internal/config"
	"github.com/IlmarLopez/currency/internal/currency"
	"github.com/IlmarLopez/currency/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Version is the version of the application.
var Version = "1.0.0"

// newServer creates the hooks and returns the router.
func newServer(lc fx.Lifecycle, db *pgx.Conn, logger *zap.SugaredLogger) *gin.Engine {
	router := gin.New()

	// build HTTP server
	address := fmt.Sprintf(":%v", 80)
	server := http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Infof("server %v is running at %v", Version, address)
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Infof("server %v is stopped", Version)
				if err := db.Close(ctx); err != nil {
					return err
				}

				return server.Shutdown(ctx)
			},
		},
	)

	return router
}

// register is the function that builds the handler.
func register(r *gin.Engine, db *pgx.Conn, logger *zap.SugaredLogger) {
	v1 := r.Group("/v1")

	// register handlers
	currency.RegisterHandlers(v1,
		currency.NewService(currency.NewRepository(db, logger), logger), logger,
	)

}

// builFlagConfig is the function that builds the flag config.
func builFlagConfig() *string {
	return flag.String("config", "./config/local.yml", "path to the configuration file")
}

func newConnectionDB(ctx context.Context, cfg *config.Config, logger *zap.SugaredLogger) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.DSN)
	if err != nil {
		logger.Errorf("error connecting to postgres: %v", err)
		return nil, err
	}

	return conn, nil
}

func main() {
	app := fx.New(
		fx.Provide(
			builFlagConfig,
			log.NewLogger,
			config.Load,
			newConnectionDB,
			newServer,
		),
		fx.Invoke(register),
	)

	app.Run()
}
