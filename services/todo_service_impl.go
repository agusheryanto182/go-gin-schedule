package services

import (
	"github.com/agusheryanto182/go-todo/helpers"
	"github.com/agusheryanto182/go-todo/models/domain"
	"github.com/agusheryanto182/go-todo/models/web"
	"github.com/agusheryanto182/go-todo/repositories"
	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepository
	validate       *validator.Validate
}

func (s *TodoServiceImpl) Create(input web.TodoCreateRequest) (web.TodoResponse, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return web.TodoResponse{}, err
	}

	todos := domain.Todo{}
	todos.ActivityGroupId = input.ActivityGroupId
	todos.Title = input.Title

	if input.Priority == "" {
		todos.Priority = "very-high"
	} else {
		todos.Priority = input.Priority
	}

	if !input.IsActive {
		todos.IsActive = true
	} else {
		todos.IsActive = input.IsActive
	}

	newTodo, err := s.todoRepository.Save(todos)
	if err != nil {
		return web.TodoResponse{}, err
	}

	return helpers.ToTodoResponse(newTodo), nil

}

func (s *TodoServiceImpl) Update(input web.TodoUpdateRequest) (web.TodoResponse, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return web.TodoResponse{}, err
	}

	todo, err := s.todoRepository.FindById(input.TodoId)
	if err != nil {
		return web.TodoResponse{}, err
	}

	grupId := todo.ActivityGroupId
	title := todo.Title
	isActive := todo.IsActive

	if input.Title == "" {
		todo.Title = title
	} else {
		todo.Title = input.Title
	}

	if input.ActivityGroupId == 0 {
		todo.ActivityGroupId = grupId
	} else {
		todo.ActivityGroupId = input.ActivityGroupId
	}

	if input.Priority == "" {
		todo.Priority = "very-high"
	} else {
		todo.Priority = input.Priority
	}

	if input.IsActive == nil {
		todo.IsActive = isActive
	} else {
		todo.IsActive = *input.IsActive
	}

	updated, err := s.todoRepository.Update(todo)
	if err != nil {
		return web.TodoResponse{}, err
	}

	return helpers.ToTodoResponse(updated), nil

}

func (s *TodoServiceImpl) Delete(Id int) error {
	err := s.todoRepository.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoServiceImpl) GetById(Id int) (web.TodoResponse, error) {
	todo, err := s.todoRepository.FindById(Id)
	if err != nil {
		return web.TodoResponse{}, err
	}

	if todo.TodoId <= 0 {
		return web.TodoResponse{}, err
	}

	return helpers.ToTodoResponse(todo), nil
}

func (s *TodoServiceImpl) GetAll() ([]domain.Todo, error) {
	todo, err := s.todoRepository.FindAll()
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *TodoServiceImpl) GetAllByActivityGroupId(activityGroupId int) ([]domain.Todo, error) {
	todo, err := s.todoRepository.FindAllByActivityGroupId(activityGroupId)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func NewTodoService(todoRepository repositories.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		todoRepository: todoRepository,
		validate:       validate,
	}
}
