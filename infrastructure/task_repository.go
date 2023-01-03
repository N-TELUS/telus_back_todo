package infrastructure

import (
	"telus_back_todo/domain"
	"telus_back_todo/domain/repository"

	"gorm.io/gorm"
)

type todoRepository struct {
	Conn *gorm.DB
}

func NewTodoRepository(conn *gorm.DB) repository.TodoRepository {
	return &todoRepository{Conn: conn}
}

func (tr *todoRepository) GetAll(userId int) ([]*domain.TodoList, error) {
	var todo []*domain.TodoList
	if err := tr.Conn.Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (tr *todoRepository) GetList(userId, listId int) (*domain.List, error) {
	list := &domain.List{UserId: userId, ListId: listId}
	if err := tr.Conn.First(&list, userId, listId).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (tr *todoRepository) GetTask(userId, taskId int) (*domain.Task, error) {
	task := &domain.Task{UserId: userId, TaskId: taskId}
	if err := tr.Conn.First(&task, userId, taskId).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *todoRepository) SaveList(list *domain.List) (*domain.List, error) {
	if err := tr.Conn.Save(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (tr *todoRepository) SaveTask(task *domain.Task) (*domain.Task, error) {
	if err := tr.Conn.Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *todoRepository) RemoveList(userId, listId int) error {
	if err := tr.Conn.Delete(userId, listId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) RemoveTask(userId, taskId int) error {
	if err := tr.Conn.Delete(userId, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) UpdateList(list *domain.List) (*domain.List, error) {
	if err := tr.Conn.Model(&list).Save(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (tr *todoRepository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	if err := tr.Conn.Model(&task).Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
