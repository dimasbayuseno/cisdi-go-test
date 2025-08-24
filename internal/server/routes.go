package server

import (
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/repository"
	"github.com/dimasbayuseno/cisdi-go-test/internal/domain/example_domain/service"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Routes() {

	api := s.app.Group("/api")
	api.Get("/health-check", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	exampleRepo := repository.New(s.db)
	exampleSvc := service.New(exampleRepo)
	exampleCtrl := example_domain.New(exampleSvc)

	s.RoutesExample(api, exampleCtrl)
}

func (s Server) RoutesExample(route fiber.Router, ctrl *example_domain.ControllerHTTP) {
	v1 := route.Group("/v1")
	exampleV1 := v1.Group("/example")
	exampleV1.Post("", ctrl.Create)
	exampleV1.Get("/:id", ctrl.GetByID)
	exampleV1.Put("/:id", ctrl.Update)
	exampleV1.Delete("/:id", ctrl.Delete)
}
