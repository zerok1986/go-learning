package usecases

import "github.com/zerok1986/go-learning/react-golang-app/server/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
