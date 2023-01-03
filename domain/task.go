package domain

import (
	"errors"
	"time"
)

type Task struct {
	UserId     int
	ListId     int
	TaskId     int
	TaskTitle  string
	NoticeDate time.Time
	CheckFlg   bool
}

type List struct {
	UserId    int
	ListId    int
	ListTitle string
	ColorId   string
}

type TodoList struct {
	ListId    int
	ListTitle string
	ColorId   string
	TaskList  Task
}

func (l *List) SetList(list *List) error {
	// Validation
	if list.ListTitle == "" {
		return errors.New("リスト名を入力してください")
	}

	l.ListTitle = list.ListTitle

	return nil
}

func (t *Task) SetTask(task *Task) error {
	// Validation
	if task.TaskTitle == "" {
		return errors.New("タスク名を入力してください")
	}

	t.ListId = task.ListId
	t.TaskTitle = task.TaskTitle
	t.NoticeDate = task.NoticeDate

	return nil
}
