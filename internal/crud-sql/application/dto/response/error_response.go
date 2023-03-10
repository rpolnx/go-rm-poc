package response_dto

import (
	"time"
)

type ErrorResponseDTO struct {
	Message   string    `json:"message,omitempty"`
	ErrorCode int       `json:"error_code,omitempty"`
	Path      string    `json:"path,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
