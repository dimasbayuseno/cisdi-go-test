package article_domain

import (
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/internal/model"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/exception"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/middleware"
	"github.com/dimasbayuseno/cisdi-go-test/pkg/pkgutil"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strconv"
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

	user := ctx.Locals("user").(middleware.UserData)
	id, err := uuid.Parse(user.ID)
	if err != nil {
		return err
	}
	req.AuthorID = id

	err = c.svc.Create(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "Article created successfully",
	})
}

// @Summary Get All Articles
// @Description Get a list of all articles with optional filtering and pagination
// @Tags Article
// @Accept json
// @Produce json
// @Param status query string false "Filter by status (draft, published)"
// @Param author_id query string false "Filter by author ID"
// @Param tag_id query string false "Filter by tag ID"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Param sort_by query string false "Sort field (default: created_at)"
// @Param sort_order query string false "Sort order (asc, desc) (default: desc)"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.GetArticlesResponse}
// @Failure 400 {object} pkgutil.HTTPResponse "Bad request"
// @Failure 500 {object} pkgutil.HTTPResponse "Internal server error"
// @Router /api/v1/article [get]
func (c ControllerHTTP) GetAll(ctx *fiber.Ctx) error {
	// Parse query parameters
	req := model.GetArticlesRequest{}

	// Parse status filter
	if status := ctx.Query("status"); status != "" {
		req.Status = status
	}

	// Parse author_id filter
	if authorIDStr := ctx.Query("author_id"); authorIDStr != "" {
		authorID, err := uuid.Parse(authorIDStr)
		if err == nil {
			req.AuthorID = authorID
		}
	}

	// Parse tag_id filter
	if tagIDStr := ctx.Query("tag_id"); tagIDStr != "" {
		tagID, err := uuid.Parse(tagIDStr)
		if err == nil {
			req.TagID = tagID
		}
	}

	// Parse pagination parameters
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	req.Page = page

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	req.Limit = limit

	// Parse sorting parameters
	if sortBy := ctx.Query("sort_by"); sortBy != "" {
		req.SortBy = sortBy
	} else {
		req.SortBy = "created_at"
	}

	if sortOrder := ctx.Query("sort_order"); sortOrder != "" {
		req.SortOrder = sortOrder
	} else {
		req.SortOrder = "desc"
	}

	// Get user information from context
	user := ctx.Locals("user").(middleware.UserData)
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return err
	}

	// Call service to get articles
	response, err := c.svc.GetArticles(ctx.UserContext(), userID, user.Role, req)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Articles retrieved successfully",
		Data:    response,
	})
}

// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.ArticleDetailResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [get]
func (c ControllerHTTP) GetBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	res, err := c.svc.GetDetailArticleBySlug(ctx.UserContext(), slug)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// @Summary Update Article
// @Description Update article by ID with status change capability
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Param status query string false "Article Status" Enums(draft,published,archived)
// @Param Authorization header string true "Bearer Token" default(<token>)
// @Success 200 {object} pkgutil.HTTPResponse{message=string} "Article updated successfully"
// @Failure 400 {object} pkgutil.HTTPResponse "Bad Request - Invalid article ID or status"
// @Failure 401 {object} pkgutil.HTTPResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} pkgutil.HTTPResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} pkgutil.HTTPResponse "Article not found"
// @Failure 500 {object} pkgutil.HTTPResponse "Internal server error"
// @Router /api/v1/articles/update/{id} [put]
func (c ControllerHTTP) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	status := ctx.Query("status")
	user := ctx.Locals("user").(middleware.UserData)

	err := c.svc.UpdateArticle(ctx.UserContext(), id, status, user.ID, user.Role)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Article updated successfully",
	})
}

// @Summary CreateNewArticleVersion Article
// @Description CreateNewArticleVersion Article
// @Tags Article
// @Accept json
// @Produce json
// @Param body body model.ArticleUpdateRequest true "Payload Article Create Request"
// @Success 201 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{errors=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/article/{id}/version [post]
func (c ControllerHTTP) CreateNewArticleVersion(ctx *fiber.Ctx) error {
	var req model.ArticleUpdateRequest
	err := ctx.BodyParser(&req)

	user := ctx.Locals("user").(middleware.UserData)
	req.AuthorID = user.ID
	req.ArticleID = ctx.Params("id")

	exception.PanicIfNeeded(err)

	err = c.svc.CreateNewArticleVersion(ctx.UserContext(), req)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "New version created successfully",
	})
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
func (c ControllerHTTP) CreateNewTag(ctx *fiber.Ctx) error {
	var req model.ArticleCreateRequest
	err := ctx.BodyParser(&req)
	exception.PanicIfNeeded(err)

	user := ctx.Locals("user").(middleware.UserData)

	id := ctx.Params("id")

	err = c.svc.CreateNewTagByAdmin(ctx.UserContext(), id, user.Role)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusCreated,
		Message: "Article created successfully",
	})
}

// @Summary Get Tag By ID
// @Description Get Tag By ID
// @Tags Tag
// @Accept json
// @Produce json
// @Success 200 {object} pkgutil.HTTPResponse{data=model.TagResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/tags [get]
func (c ControllerHTTP) GetAllTag(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(middleware.UserData)
	res, err := c.svc.GetAllTags(ctx.UserContext(), user.Role)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// @Summary Get Tag By ID
// @Description Get Tag By ID
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.TagResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/tags/{id} [get]
func (c ControllerHTTP) GetTagByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user").(middleware.UserData)
	res, err := c.svc.GetDetailTag(ctx.UserContext(), id, user.Role)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// @Summary Delete Article
// @Description Delete Article
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/articles/{id} [delete]
func (c ControllerHTTP) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user").(middleware.UserData)
	err := c.svc.DeleteArticle(ctx.UserContext(), id, user.ID, user.Role)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "User deleted successfully",
	})
}

// @Summary Get All Version Article By ID
// @Description Get Article By ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.ArticleDetailResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [get]
func (c ControllerHTTP) GetAllArticleVersion(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user").(middleware.UserData)
	idNew, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}
	res, err := c.svc.GetArticleVersions(ctx.UserContext(), idNew, user.Role, user.ID)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}

// @Summary Get All Version Article By ID
// @Description Get Article By ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} pkgutil.HTTPResponse{data=model.ArticleDetailResponse} "Success"
// @Failure 404 {object} pkgutil.HTTPResponse
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /api/v1/users/{id} [get]
func (c ControllerHTTP) GetSpecificArticleVersion(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	versionNumber := ctx.Query("version")
	version, err := strconv.Atoi(versionNumber)
	if err != nil {
		return fmt.Errorf(": %v", err)
	}
	res, err := c.svc.GetDetailArticleVersion(ctx.UserContext(), id, version)
	exception.PanicIfNeeded(err)

	return ctx.JSON(pkgutil.HTTPResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    res,
	})
}
