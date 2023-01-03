package usecase

import (
	"telus_back_todo/domain"
	"telus_back_todo/domain/repository"
)

type TodoUsecase interface {
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

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: todoRepo}
}

func (tu *todoUsecase) GetAll(userId int) ([]*domain.TodoList, error) {
	foundTodo, err := tu.todoRepo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	return foundTodo, nil
}

func (tu *todoUsecase) GetList(userId, listId int) (*domain.List, error) {
	foundList, err := tu.todoRepo.GetList(userId, listId)
	if err != nil {
		return nil, err
	}

	return foundList, nil
}

func (tu *todoUsecase) GetTask(userId, taskId int) (*domain.Task, error) {
	foundTask, err := tu.todoRepo.GetTask(userId, taskId)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *todoUsecase) SaveList(list *domain.List) (*domain.List, error) {
	err := list.SetList(list)
	if err != nil {
		return nil, err
	}

	saveTask, err := tu.todoRepo.SaveList(list)
	if err != nil {
		return nil, err
	}

	return saveTask, nil
}

func (tu *todoUsecase) SaveTask(task *domain.Task) (*domain.Task, error) {
	err := task.SetTask(task)
	if err != nil {
		return nil, err
	}

	saveTask, err := tu.todoRepo.SaveTask(task)
	if err != nil {
		return nil, err
	}

	return saveTask, nil
}

func (tu *todoUsecase) RemoveList(userId, listId int) error {
	list, err := tu.todoRepo.GetList(userId, listId)
	if err != nil {
		return err
	}

	err = tu.todoRepo.RemoveList(list.UserId, list.ListId)
	if err != nil {
		return err
	}

	return nil
}

func (tu *todoUsecase) RemoveTask(userId, taskId int) error {
	task, err := tu.todoRepo.GetTask(userId, taskId)
	if err != nil {
		return err
	}

	err = tu.todoRepo.RemoveTask(task.UserId, task.TaskId)
	if err != nil {
		return err
	}

	return nil
}

func (tu *todoUsecase) UpdateList(list *domain.List) (*domain.List, error) {
	targetList, err := tu.todoRepo.GetList(list.UserId, list.ListId)
	if err != nil {
		return nil, err
	}

	err = targetList.SetList(list)
	if err != nil {
		return nil, err
	}

	updatedList, err := tu.todoRepo.SaveList(targetList)
	if err != nil {
		return nil, err
	}

	return updatedList, err
}

func (tu *todoUsecase) UpdateTask(task *domain.Task) (*domain.Task, error) {
	targetTask, err := tu.todoRepo.GetTask(task.UserId, task.TaskId)
	if err != nil {
		return nil, err
	}

	err = targetTask.SetTask(task)
	if err != nil {
		return nil, err
	}

	updatedTask, err := tu.todoRepo.SaveTask(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, err
}
