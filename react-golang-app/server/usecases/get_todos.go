package usecases

import "github.com/zerok1986/go-learning/react-golang-app/server/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
	return todos, nil
}
