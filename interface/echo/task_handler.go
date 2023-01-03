package echo

import (
	"net/http"
	"strconv"
	"time"

	"telus_back_todo/domain"
	"telus_back_todo/usecase"

	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	PostList() echo.HandlerFunc
	PostTask() echo.HandlerFunc
	GetList() echo.HandlerFunc
	GetTask() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	PutList() echo.HandlerFunc
	PutTask() echo.HandlerFunc
	DeleteList() echo.HandlerFunc
	DeleteTask() echo.HandlerFunc
}

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	return &todoHandler{todoUsecase: todoUsecase}
}

type requestList struct {
	UserId    int    `json:"userId"`
	ListTitle string `json:"listTitle"`
	ColorId   string `json:"colorId"`
}

type requestTask struct {
	UserId     int       `json:"userId"`
	ListId     int       `json:"listId"`
	TaskTitle  string    `json:"taskTitle"`
	NoticeDate time.Time `json:"noticeDate"`
	CheckFlg   bool      `json:"checkFlg"`
}

type responseList struct {
	ListId    int    `json:"listId"`
	ListTitle string `json:"listTitle"`
	ColorId   string `json:"colorId"`
}

type responseTask struct {
	ListId     int       `json:"listId"`
	TaskTitle  string    `json:"taskTitle"`
	NoticeDate time.Time `json:"noticeDate"`
	CheckFlg   bool      `json:"checkFlg"`
}

type responseTodo struct {
	ListId    int         `json:"listId"`
	ListTitle string      `json:"listTitle"`
	ColorId   string      `json:"colorId"`
	TaskList  domain.Task `json:"taskList"`
}

func (th *todoHandler) PostList() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestList
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdList, err := th.todoUsecase.SaveList(
			&domain.List{
				UserId:    req.UserId,
				ListTitle: req.ListTitle,
				ColorId:   req.ColorId,
			})

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := requestList{
			UserId:    createdList.UserId,
			ListTitle: createdList.ListTitle,
			ColorId:   createdList.ColorId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (th *todoHandler) PostTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.todoUsecase.SaveTask(
			&domain.Task{
				UserId:     req.UserId,
				ListId:     req.ListId,
				TaskTitle:  req.TaskTitle,
				NoticeDate: req.NoticeDate,
			})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := requestTask{
			UserId:     createdTask.UserId,
			ListId:     createdTask.ListId,
			TaskTitle:  createdTask.TaskTitle,
			NoticeDate: createdTask.NoticeDate,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (th *todoHandler) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		listId, err := strconv.Atoi((c.Param("listId")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundList, err := th.todoUsecase.GetList(1, listId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := requestList{
			UserId:    foundList.UserId,
			ListTitle: foundList.ListTitle,
			ColorId:   foundList.ColorId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *todoHandler) GetTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, err := strconv.Atoi((c.Param("taskId")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.todoUsecase.GetTask(1, taskId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := requestTask{
			UserId:     foundTask.UserId,
			ListId:     foundTask.ListId,
			TaskTitle:  foundTask.TaskTitle,
			NoticeDate: foundTask.NoticeDate,
			CheckFlg:   foundTask.CheckFlg,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *todoHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi((c.Param("userId")))

		foundTodo, err := th.todoUsecase.GetAll(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseTodo{}
		for _, todo := range foundTodo {
			res = append(res, responseTodo{
				ListId:    todo.ListId,
				ListTitle: todo.ListTitle,
				ColorId:   todo.ColorId,
				TaskList:  todo.TaskList,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *todoHandler) PutList() echo.HandlerFunc {
	return func(c echo.Context) error {
		listId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestList
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedList, err := th.todoUsecase.UpdateList(&domain.List{
			UserId: 1, ListId: listId, ListTitle: req.ListTitle, ColorId: req.ColorId,
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseList{
			ListId:    updatedList.ListId,
			ListTitle: updatedList.ListTitle,
			ColorId:   updatedList.ColorId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *todoHandler) PutTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.todoUsecase.UpdateTask(&domain.Task{
			UserId: 1, ListId: taskId, TaskId: taskId, TaskTitle: req.TaskTitle, NoticeDate: req.NoticeDate, CheckFlg: req.CheckFlg,
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ListId:     updatedTask.ListId,
			TaskTitle:  updatedTask.TaskTitle,
			NoticeDate: updatedTask.NoticeDate,
			CheckFlg:   updatedTask.CheckFlg,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *todoHandler) DeleteList() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId取得関数

		listId, err := strconv.Atoi(c.Param("listId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.todoUsecase.RemoveList(1, listId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func (th *todoHandler) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userId取得関数

		taskId, err := strconv.Atoi(c.Param("taskId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.todoUsecase.RemoveTask(1, taskId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
