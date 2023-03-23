package postgres

import (
	"regexp"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"strings"
)

func HandleDbError(dbError error) error {

	if dbError == nil {
		return nil
	}

	message := dbError.Error()
	switch {
	case strings.Contains(message, "no rows in result set"):
		return &config.NotFoundError{Name: "User"}
	case strings.Contains(message, "ERROR #"):
		re := regexp.MustCompile("ERROR #(\\d+)\\s(.*)")

		res := re.FindAllStringSubmatch(message, 1)

		return &config.DbError{
			ErrorCode:   res[0][1],
			Cause:       res[0][2],
			FullMessage: message,
		}
	case strings.Contains(message, "sql: "):
		re := regexp.MustCompile("sql:\\s(.*)")

		res := re.FindAllStringSubmatch(message, 1)

		return &config.DbError{
			ErrorCode:   "SQl_DB_ERROR",
			Cause:       res[0][1],
			FullMessage: message,
		}
	default:
		return dbError
	}
}
