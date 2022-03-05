package ui

import "github.com/zerok1986/go-learning/react-golang-app/server/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
