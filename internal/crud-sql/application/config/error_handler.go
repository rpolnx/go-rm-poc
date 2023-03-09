package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type ErrorResponse struct {
	Message   string    `json:"message,omitempty"`
	ErrorCode int       `json:"error_code,omitempty"`
	Path      string    `json:"path,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func HandleError(ginCtx *gin.Context, statusCode int, err error) *ErrorResponse {

	response := ErrorResponse{
		Message:   err.Error(),
		ErrorCode: statusCode,
		Path:      ginCtx.FullPath(),
		Timestamp: time.Now(),
	}

	logrus.Error(response)

	return &response
}
