package server

import (
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/service"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain"
	user_repository "github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain/repository"
	user_service "github.com/dimasbayuseno/cisdi-go-test/internal/domain/user_domain/service"
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

	s.RoutesExample(api, exampleCtrl)
	s.RoutesUser(api, userCtrl)
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
	userV1.Get("/login", ctrl.Login)
	userV1.Get("/:id", ctrl.GetByID)
	userV1.Put("/:id", ctrl.Update)
	userV1.Delete("/:id", ctrl.Delete)
}
