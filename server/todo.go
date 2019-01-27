package server

import (
	"github.com/chelium/simple-website/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todoHandler struct {
	s todoService
}

func (h *todoHandler) AddRoutes(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/", getTodosEndpoint)
		todos.GET("/:id", getTodoEndpoint)
		todos.PUT("/:id", putTodoEndpoint)
		todos.POST("/", postTodoEndpoint)
		todos.DELETE("/:id", deleteTodoEndpoint)
	}
}

func getTodosEndpoint(c *gin.Context) {
	userID := c.MustGet(gin.AuthUserKey).(string)
	getTodos, err := s.GetUserTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodos)
}

func getTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	getTodo, err := s.GetUserTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}

func postTodoEndpoint(c *gin.Context) {
	userID := c.MustGet(gin.AuthUserKey).(string)
	newTodo := todo.Todo{}
	if err := c.ShouldBindBodyWith(newTodo, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newT := todo.NewTodo(newTodo.Title, newTodo.Description)
	newTID, err := s.CreateUserTodo(userID, newT)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newTID)
}

func putTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	t := todo.Todo{}
	if err := c.ShouldBindBodyWith(t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := s.UpdateUserTodo(userID, todoID, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newTID)
}

func deleteTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	err := s.DeleteUserTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK)
}
