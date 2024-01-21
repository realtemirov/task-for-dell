package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/docs"
	blogHttpV1 "github.com/realtemirov/task-for-dell/internal/blogs/delivery/http"
	blogRepo "github.com/realtemirov/task-for-dell/internal/blogs/repository"
	blogUseCase "github.com/realtemirov/task-for-dell/internal/blogs/usecase"

	newsHttpV1 "github.com/realtemirov/task-for-dell/internal/news/delivery/http"
	newsRepo "github.com/realtemirov/task-for-dell/internal/news/repository"
	newsUseCase "github.com/realtemirov/task-for-dell/internal/news/usecase"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	MAX_HEADER_SIZE = 1 << 20 // 1 MB
	STACK_SIZE      = 1 << 10 // 1 KB
	BODY_LIMIT      = "2M"
	GZIP_LEVEL      = 5
)

type server struct {
	cfg     *config.Config
	log     logger.Logger
	psql    *sqlx.DB
	echo    *echo.Echo
	validat *validator.Validate
}

func NewServer(cfg *config.Config, log logger.Logger, psql *sqlx.DB) *server {
	return &server{
		cfg:     cfg,
		log:     log,
		psql:    psql,
		echo:    echo.New(),
		validat: validator.New(),
	}
}

func (s *server) Run() error {

	// init server
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: MAX_HEADER_SIZE,
		Handler:        s.echo,
	}

	go func() {
		s.log.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.log.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), s.cfg.Server.CtxDefaultTime*time.Second)
	defer shutdown()

	s.log.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}

// @Summary Health check endpoint
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} model "{"message":"pong"}"
// @Router /ping [get]
// Map Server Handlers
func (s *server) MapHandlers(e *echo.Echo) error {

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Blog and News API."
	docs.SwaggerInfo.Description = "Blog and News API Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         STACK_SIZE,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: GZIP_LEVEL,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit(BODY_LIMIT))

	v1 := s.echo.Group("/v1")

	// blogs
	blogPGRepo := blogRepo.NewBlogsRepository(s.psql)
	blogUC := blogUseCase.NewBlogUseCase(s.cfg, blogPGRepo, s.log)
	blogHandler := blogHttpV1.NewBlogsHandlers(s.cfg, blogUC, s.log)
	blogHttpV1.MapBlogsRoutes(v1.Group("/blogs"), blogHandler)

	// news
	newsPGRepo := newsRepo.NewNewsRepository(s.psql)
	newsUC := newsUseCase.NewNewsUseCase(s.cfg, newsPGRepo, s.log)
	newsHandler := newsHttpV1.NewNewsHandlers(s.cfg, newsUC, s.log)
	newsHttpV1.MapNewsRoutes(v1.Group("/news"), newsHandler)

	v1.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	return nil
}
