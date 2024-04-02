package http

import (
	"errors"
	"net/http"
	"todo/internal/pkg/app_err"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResponse(c *gin.Context, err error) {
	var bErr app_err.BusinessError

	if errors.As(err, &bErr) {
		errorResponse := ErrorResponse{
			Error: Error{
				Code:    bErr.Code(),
				Message: bErr.Error(),
			},
		}

		c.JSON(http.StatusBadRequest, errorResponse)
	} else {
		errorResponse := ErrorResponse{
			Error: Error{
				Code:    "InternalServerError",
				Message: "Что-то пошло не так, попробуйте еще раз",
			},
		}

		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
