package config

import "fmt"

type NotFoundError struct {
	Name string
}

func (n DbError) Error() string {
	return "Entity not found"
}

type DbError struct{}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("Entity %s not found", n.Name)
}
