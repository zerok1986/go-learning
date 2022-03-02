package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/zerok1986/go-learning/react-golang-app/server/entities"
	"github.com/zerok1986/go-learning/react-golang-app/server/usecases"
)

var dummyTodos = []entities.Todo{
	{
		Title:       "todo 1",
		Description: "description of todo1",
		IsCompleted: true,
	},
	{
		Title:       "todo 2",
		Description: "description of todo2",
		IsCompleted: false,
	},
	{
		Title:       "todo 3",
		Description: "description of todo3",
		IsCompleted: false,
	},
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("Something went wrong")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	// Test
	t.Run("Returns ErrInternal when TodosRepository returns error", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(BadTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)
		if todos != nil {
			t.Fatalf("Expected todos to be nil; Got: %v", todos)
		}
	})

	// Test2
	t.Run("Returns todos from TodosRepository", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}
