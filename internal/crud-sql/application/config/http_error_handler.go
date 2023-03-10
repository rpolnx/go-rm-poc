package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/dto/response"
	"time"
)

// HandleHttpError handle generic application errors
// Provide a status_code as optional field on third param
func HandleHttpError(ginCtx *gin.Context, err error, params ...int) (int, *response_dto.ErrorResponseDTO) {
	httpErrorCode := 999

	if len(params) > 0 {
		httpErrorCode = params[0]
	} else {
		httpErrorCode = getHttpStatusCodeByError(err)
	}

	value := response_dto.ErrorResponseDTO{
		Message:   err.Error(),
		ErrorCode: httpErrorCode,
		Path:      ginCtx.FullPath(),
		Timestamp: time.Now(),
	}

	logrus.Error(value)

	return httpErrorCode, &value
}

func getHttpStatusCodeByError(err error) int {
	switch err.(type) {
	case *NotFoundError:
		return http.StatusNotFound
	case *DbError:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}

}
