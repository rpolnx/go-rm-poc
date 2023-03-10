package config

import "fmt"

type NotFoundError struct {
	Name string
}

func (n DbError) Error() string {
	return fmt.Sprintf("Error #%s processing entity with cause %s", n.ErrorCode, n.Cause)
}

type DbError struct {
	ErrorCode   string
	Cause       string
	FullMessage string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("Entity %s not found", n.Name)
}
