package helper

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginationResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Pagination  `json:"meta"`
}

func Success(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, Response{
		Success: false,
		Message: message,
	})
}

func SuccessWithPagination(ctx *gin.Context, statusCode int, message string, data interface{}, meta Pagination) {
	ctx.JSON(statusCode, PaginationResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
