package server

import (
	"github.com/chelium/simple-website/todo"
	ts "github.com/chelium/simple-website/todo/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type todoHandler struct {
	s ts.Service
}

func (h *todoHandler) AddRoutes(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/", h.getTodosEndpoint)
		todos.GET("/:id", h.getTodoEndpoint)
		todos.PUT("/:id", h.putTodoEndpoint)
		todos.POST("/", h.postTodoEndpoint)
		todos.DELETE("/:id", h.deleteTodoEndpoint)
	}
}

func (h *todoHandler) getTodosEndpoint(c *gin.Context) {
	userID := c.MustGet(gin.AuthUserKey).(string)
	getTodos, err := h.s.GetUserTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodos)
}

func (h *todoHandler) getTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	getTodo, err := h.s.GetUserTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}

func (h *todoHandler) postTodoEndpoint(c *gin.Context) {
	userID := c.MustGet(gin.AuthUserKey).(string)
	newTodo := todo.Todo{}
	if err := c.ShouldBindBodyWith(newTodo, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newT := todo.NewTodo(newTodo.Title, newTodo.Description)
	newTID, err := h.s.CreateUserTodo(userID, *newT)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newTID)
}

func (h *todoHandler) putTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	t := todo.Todo{}
	if err := c.ShouldBindBodyWith(t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.s.UpdateUserTodo(userID, todoID, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *todoHandler) deleteTodoEndpoint(c *gin.Context) {
	todoID := c.Param("id")
	userID := c.MustGet(gin.AuthUserKey).(string)
	err := h.s.DeleteUserTodo(userID, todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
