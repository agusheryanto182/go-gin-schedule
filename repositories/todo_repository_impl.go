package repositories

import (
	"github.com/agusheryanto182/go-todo/models/domain"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func (r *TodoRepositoryImpl) Save(todo domain.Todo) (domain.Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Update(todo domain.Todo) (domain.Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Delete(Id int) error {
	err := r.db.Delete(&domain.Todo{}, "todo_id = ?", Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepositoryImpl) FindById(Id int) (domain.Todo, error) {
	var todo domain.Todo
	err := r.db.Where("todo_id = ?", Id).Find(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) FindAll() ([]domain.Todo, error) {
	var todo []domain.Todo
	err := r.db.Find(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) FindAllByActivityGroupId(activityGroupId int) ([]domain.Todo, error) {
	var todo []domain.Todo
	err := r.db.Where("activity_group_id = ?", activityGroupId).Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}
