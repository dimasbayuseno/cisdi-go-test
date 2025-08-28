package pkgutil

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Code    int    `json:"code" example:"200"`
	Status  string `json:"status" example:"OK"`
	Message string `json:"message,omitempty" example:"Success"`
	Data    any    `json:"data,omitempty" `
	Errors  []any  `json:"errors,omitempty" `
}

func (h HTTPResponse) MarshalJSON() ([]byte, error) {
	type Alias HTTPResponse
	alias := &struct {
		*Alias
	}{
		Alias: (*Alias)(&h),
	}

	alias.Status = http.StatusText(h.Code)
	if alias.Status == alias.Message {
		alias.Message = ""
	}

	return json.Marshal(alias)
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

type ErrValidationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func BuildPagination(page, limit int, totalCount int64) PaginationResponse {
	if limit <= 0 {
		limit = 10
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	hasPrev := page > 1
	hasNext := page < totalPages

	return PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalCount: totalCount,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
}
