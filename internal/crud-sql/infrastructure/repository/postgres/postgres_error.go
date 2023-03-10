package postgres

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"strings"
)

func HandleDbError(dbError error) error {

	if dbError == nil {
		return dbError
	}

	message := dbError.Error()
	switch {
	case strings.Contains(message, "pg: no rows in result set"):
		return &config.NotFoundError{Name: "User"}
	case strings.Contains(message, "test"):
		return &config.DbError{}
	default:
		return dbError
	}
}
