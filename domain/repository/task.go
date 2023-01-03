package repository

import (
	"telus_back_todo/domain"
)

type TodoRepository interface {
	GetAll(userId int) ([]*domain.TodoList, error)
	GetList(userId, listId int) (*domain.List, error)
	GetTask(userId, taskId int) (*domain.Task, error)
	SaveList(task *domain.List) (*domain.List, error)
	SaveTask(task *domain.Task) (*domain.Task, error)
	RemoveList(userId, listId int) error
	RemoveTask(userId, taskId int) error
	UpdateList(task *domain.List) (*domain.List, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
}
