package example_domain

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

// @Summary Create Example
// @Description Create Example
// @Tags Example
// @Accept json
// @Produce json
// @Param body body model.ExampleCreateRequest true "Payload Example Create Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/examples [post]
func (c ControllerHTTP) Create(ctx *fiber.Ctx) error {
	var req model.ExampleCreateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = c.svc.Create(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "Example created successfully",
	})
}

// @Summary Get Example By ID
// @Description Get Example By ID
// @Tags Example
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.ExampleResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/examples/{id} [get]
func (c ControllerHTTP) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := c.svc.GetByID(ctx.UserContext(), id)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// @Summary Update Example
// @Description Update Example
// @Tags Example
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Param body body model.ExampleUpdateRequest true "Payload Example Update Request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/examples/{id} [put]
func (c ControllerHTTP) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req model.ExampleUpdateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	req.ID = id
	err = c.svc.Update(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Example updated successfully",
	})
}

// @Summary Delete Example
// @Description Delete Example
// @Tags Example
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/examples/{id} [delete]
func (c ControllerHTTP) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.svc.Delete(ctx.UserContext(), id)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Example deleted successfully",
	})
}
