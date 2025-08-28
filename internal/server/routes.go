package server

import (
	"github.com/dimasbayuseno/cisdi-go-test/config"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain"
	article_repository "github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/repository"
	article_service "github.com/dimasbayuseno/cisdi-go-test/internal/domain/article_domain/service"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/service"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain"
	user_repository "github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain/repository"
	user_service "github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain/service"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Routes() {

	api := s.app.Group("/api")
	api.Get("/health-check", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	exampleRepo := repository.New(s.db)
	exampleSvc := service.New(exampleRepo)
	exampleCtrl := example_domain.New(exampleSvc)

	userRepo := user_repository.New(s.db)
	userSvc := user_service.New(userRepo)
	userCtrl := user_domain.New(userSvc)

	articleRepo := article_repository.New(s.db)
	articleSvc := article_service.New(articleRepo)
	articleCtrl := article_domain.New(articleSvc)

	s.RoutesExample(api, exampleCtrl)
	s.RoutesUser(api, userCtrl)
	s.RoutesArticle(api, articleCtrl)
}

func (s Server) RoutesExample(route fiber.Router, ctrl *example_domain.ControllerHTTP) {
	v1 := route.Group("/v1")
	exampleV1 := v1.Group("/example")
	exampleV1.Post("", ctrl.Create)
	exampleV1.Get("/:id", ctrl.GetByID)
	exampleV1.Put("/:id", ctrl.Update)
	exampleV1.Delete("/:id", ctrl.Delete)
}

func (s Server) RoutesUser(route fiber.Router, ctrl *user_domain.ControllerHTTP) {
	v1 := route.Group("/v1")
	userV1 := v1.Group("/user")
	userV1.Post("/register", ctrl.Create)
	userV1.Post("/login", ctrl.Login)
	userV1.Get("/:id", ctrl.GetByID)
	userV1.Put("/:id", ctrl.Update)
	userV1.Delete("/:id", ctrl.Delete)
}

func (s Server) RoutesArticle(route fiber.Router, ctrl *article_domain.ControllerHTTP) {
	v1 := route.Group("/v1")
	articleV1 := v1.Group("/article")
	articleV1.Post("/create", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.Create)
	articleV1.Post("/create-version/:id", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.CreateNewArticleVersion)
	articleV1.Get("", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetAll)
	articleV1.Put("/update/:id", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.Update)
	articleV1.Get("/:slug", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetBySlug)
	articleV1.Get("/:id/:version", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetSpecificArticleVersion)
	articleV1.Get("/all-version/:id", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetAllArticleVersion)
	articleV1.Get("/:slug", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetBySlug)
	articleV1.Delete("/delete/:id", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.Delete)

	tagV1 := v1.Group("/tag")
	tagV1.Post("/create", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.CreateNewTag)
	tagV1.Get("/:id", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetTagByID)
	tagV1.Get("", middleware.JWTMiddleware(config.Get().JwtSecret), ctrl.GetAllTag)
}
