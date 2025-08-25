package constant

import (
	"errors"
	"net/http"
)

const (
	ErrSQLUniqueViolation = "23505"
	ErrSQLInvalidUUID     = "22P02"
	ErrSQLFKViolation     = "23503"
)

var (
	ErrEmailAlreadyRegistered = &ErrWithCode{HTTPStatusCode: http.StatusConflict, Message: "email already registered"}
	ErrEmailOrPasswordInvalid = &ErrWithCode{HTTPStatusCode: http.StatusBadRequest, Message: "email or password invalid"}
	ErrInvalidUUID            = errors.New("invalid uuid length or format")
	ErrUnauthorizedAccess     = &ErrWithCode{HTTPStatusCode: http.StatusUnauthorized, Message: "unauthorized access"}
	ErrExampleNotFound        = &ErrWithCode{HTTPStatusCode: http.StatusNotFound, Message: "example not found"}
	ErrUserNotFound           = &ErrWithCode{HTTPStatusCode: http.StatusNotFound, Message: "user not found"}
	ErrTagNotFound            = &ErrWithCode{HTTPStatusCode: http.StatusNotFound, Message: "tag not found"}
	ErrArticleNotFound        = &ErrWithCode{HTTPStatusCode: http.StatusNotFound, Message: "article not found"}
	ErrFailedTx               = &ErrWithCode{HTTPStatusCode: http.StatusPreconditionFailed, Message: "tag not found"}
	ErrInvalidStatusArticle   = &ErrWithCode{HTTPStatusCode: http.StatusPreconditionFailed, Message: "status invalid"}
)

type ErrWithCode struct {
	HTTPStatusCode int
	Message        string
}

func (e *ErrWithCode) Error() string {
	return e.Message
}

type ErrValidation struct {
	Message string
}

func (e *ErrValidation) Error() string {
	return e.Message
}
