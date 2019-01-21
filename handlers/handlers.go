package handlers

import (
	"github.com/chelium/simple-website/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTodoHandler will get a specified user based on user id
func GetTodoHandler(c *gin.Context) {
	// todoId := c.Param("id")
	getTodo, err := todo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}
