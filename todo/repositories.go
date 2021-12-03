package todo

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TodoReposity struct {
	database *gorm.DB
}

func (repository *TodoReposity) FindAll() []Todo {
	var todos []Todo
	repository.database.Find(&todos)
	return todos
}

func (repository *TodoReposity) Find(id int) (Todo, error) {
	var todo Todo
	err := repository.database.Find(&todo, id).Error
	if todo.Name == "" {
		err = errors.New("Todo not found")
	}
	return todo, err
}

func (repository *TodoReposity) Create(todo Todo) (Todo, error) {
	err := repository.database.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (repository *TodoReposity) Save(user Todo) (Todo, error) {
	err := repository.database.Save(user).Error
	return user, err
}

func (repository *TodoReposity) Delete(id int) int64 {
	count := repository.database.Delete(&Todo{}, id).RowsAffected
	return count
}

func NewTodoRepository(database *gorm.DB) *TodoReposity {
	return &TodoReposity{
		database: database,
	}
}
