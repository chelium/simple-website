package server

import (
	"github.com/chelium/simple-website/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todoHandler struct {
}

func (h *todoHandler) AddRoutes(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/", getTodoEndpoint)
		todos.PUT("/:id", putTodoEndpoint)
		todos.POST("/", postTodoEndpoint)
		todos.DELETE("/:id", deleteTodoEndpoint)
	}
}

// GetTodoEndpoint will get a specified user based on user id
func getTodoEndpoint(c *gin.Context) {
	// todoId := c.Param("id")
	getTodo, err := todo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}

// PostTodoEndpoint will create a user using payload
func PostTodoEndpoint(c *gin.Context) {

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}
