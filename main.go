package main

import (
	"telus_back_todo/config"
	"telus_back_todo/infrastructure"
	"telus_back_todo/interface/echo"
	"telus_back_todo/usecase"
)

func main() {

	// repository
	todoRepository := infrastructure.NewTodoRepository(config.NewDB())
	// usecase
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	// handler
	todoHandler := echo.NewTodoHandler(todoUsecase)

	// Start server
	echo.Start(todoHandler)
}
