package article_domain

import (
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/exception"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
)

type ControllerHTTP struct {
	svc Service
}

func New(svc Service) *ControllerHTTP {
	return &ControllerHTTP{svc: svc}
}

// @Summary Create Article
// @Description Create Article
// @Tags Article
// @Accept json
// @Produce json
// @Param body body model.ArticleCreateRequest true "Payload Article Create Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/articles [post]
func (c ControllerHTTP) Create(ctx *fiber.Ctx) error {
	var req model.ArticleCreateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = c.svc.Create(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "Article created successfully",
	})
}
