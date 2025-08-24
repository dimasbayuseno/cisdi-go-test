package user_domain

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

// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param body body model.UserCreateRequest true "Payload User Create Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users [post]
func (c ControllerHTTP) Create(ctx *fiber.Ctx) error {
	var req model.UserCreateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	err = c.svc.Create(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "User created successfully",
	})
}

// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.UserResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [get]
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

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body model.UserUpdateRequest true "Payload User Update Request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [put]
func (c ControllerHTTP) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req model.UserUpdateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	req.ID = id
	err = c.svc.Update(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "User updated successfully",
	})
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [delete]
func (c ControllerHTTP) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.svc.Delete(ctx.UserContext(), id)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "User deleted successfully",
	})
}

// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "Payload User Login Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users [post]
func (c ControllerHTTP) Login(ctx *fiber.Ctx) error {
	var req model.LoginRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	res, err := c.svc.Login(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}
