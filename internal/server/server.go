package server

import (
	"context"
	"github.com/dimasbayuseno/cisdi-go-test/config"
	_ "github.com/dimasbayuseno/cisdi-go-test/docs"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/exception"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/logger"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/middleware"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

const (
	ctxTimeout = 5
)

type Server struct {
	app *fiber.App
	db  *pgxpool.Pool
}

func New(
	db *pgxpool.Pool,
) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.FiberErrorHandler,
	})
	timeout := time.Duration(config.Get().Service.Timeout) * time.Second
	app.Use(middleware.Timeout(timeout))

	app.Use(cors.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		// Logger: logger.Log(context.Background()),
		GetLogger: func(c *fiber.Ctx) zerolog.Logger {
			return *logger.Log(c.UserContext())
		},
	}))

	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Run() error {
	s.Routes()
	ctx := context.Background()
	go func() {
		if err := s.app.Listen(pkgutil.GetPort()); err != nil {
			logger.Log(ctx).Fatal().Err(err).Msg("failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(ctx, ctxTimeout*time.Second)
	defer shutdown()

	logger.Log(ctx).Info().Msg("shutting down server")
	return s.app.Shutdown()
}
