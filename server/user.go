package server

import (
	"github.com/chelium/simple-website/user"
	us "github.com/chelium/simple-website/user/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type userHandler struct {
	s us.Service
}

func (h *userHandler) AddRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/:id", h.getUserEndpoint)
		users.PUT("/:id", h.putUserEndpoint)
		users.POST("/", h.postUserEndpoint)
		users.DELETE("/:id", h.deleteUserEndpoint)
	}
}

func (h *userHandler) getUserEndpoint(c *gin.Context) {
	userID := c.Param("id")
	authUserID := c.MustGet(gin.AuthUserKey).(string)
	if userID != authUserID {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	getUser, err := h.s.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getUser)
}

func (h *userHandler) putUserEndpoint(c *gin.Context) {
	userID := c.Param("id")
	authUserID := c.MustGet(gin.AuthUserKey).(string)
	if userID != authUserID {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	updateUser := user.User{}
	if err := c.ShouldBindBodyWith(updateUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.s.UpdateUser(userID, updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *userHandler) postUserEndpoint(c *gin.Context) {
	// authUserID := c.MustGet(gin.AuthUserKey).(string)
	newUser := user.User{}
	if err := c.ShouldBindBodyWith(newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newU := user.NewUser(newUser.Username, newUser.PasswordHash)
	newUID, err := h.s.CreateUser(*newU)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newUID)
}

func (h *userHandler) deleteUserEndpoint(c *gin.Context) {
	userID := c.Param("id")
	authUserID := c.MustGet(gin.AuthUserKey).(string)
	if userID != authUserID {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	if err := h.s.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
