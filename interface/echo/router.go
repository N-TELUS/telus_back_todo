package echo

import (
	"telus_back_todo/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(todoHandler TodoHandler) {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todo", todoHandler.GetAll())
	e.GET("/list/:id", todoHandler.GetList())
	e.GET("/task/:id", todoHandler.GetTask())
	e.POST("/list", todoHandler.PostList())
	e.POST("/task", todoHandler.PostTask())
	e.PUT("/list/:id", todoHandler.PutList())
	e.PUT("/task/:id", todoHandler.PutTask())
	e.DELETE("/list/:id", todoHandler.DeleteList())
	e.DELETE("/task/:id", todoHandler.DeleteTask())

	// Start server
	e.Logger.Fatal(e.Start(":" + string(rune(config.Config.ServerPort))))
}
