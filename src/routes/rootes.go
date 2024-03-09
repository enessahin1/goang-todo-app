package routes

import (
	todo_controller "todo_app/src/todo/controllers"
	user_controller "todo_app/src/user/controllers"
	auth_middleware "todo_app/src/user/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	// Todo routes
	route.GET("/todos", todo_controller.GetAllTodos)
	route.GET("/todos/:id", todo_controller.RetriveTodo)
	route.POST("/todos", auth_middleware.AuthMiddleware(), todo_controller.CreateTodo)
	route.PATCH("/todos/:id", auth_middleware.AuthMiddleware(), todo_controller.UpdateTodo)
	route.DELETE("/todos/:id", auth_middleware.AuthMiddleware(), todo_controller.DeleteTodo)

	// User routes
	route.POST("user/login", user_controller.Login)
	route.POST("user/register", user_controller.Register)

	route.Run()
}
