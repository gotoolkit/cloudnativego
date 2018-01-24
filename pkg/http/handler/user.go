package handler

import (
	"net/http"
	"strconv"

	"github.com/gotoolkit/cloudnativego/pkg/cloudnativego"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService cloudnativego.UserService
}

func NewUserHandler(rg *gin.RouterGroup) *UserHandler {
	h := &UserHandler{}

	rg.POST("/", h.createUser)
	rg.GET("/", h.fetchAllUser)
	rg.GET("/:id", h.fetchSingleUser)
	rg.PUT("/:id", h.updateUser)
	rg.DELETE("/:id", h.deleteUser)

	return h
}

// createUser add a new User
func (handler *UserHandler) createUser(c *gin.Context) {
	var user cloudnativego.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.UserService.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User item created successfully!", "resourceId": user.ID})
}

// fetchAllUser fetch all Users
func (handler *UserHandler) fetchAllUser(c *gin.Context) {
	var users []cloudnativego.User

	users, err := handler.UserService.Users()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(users) <= 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// fetchSingleUser fetch a single User
func (handler *UserHandler) fetchSingleUser(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("id"))
	user, err := handler.UserService.User(cloudnativego.UserID(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// updateUser update a User
func (handler *UserHandler) updateUser(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("id"))
	user, err := handler.UserService.User(cloudnativego.UserID(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User found!"})
		return
	}
	var inputUser cloudnativego.User
	err = c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.UserService.UpdateUser(user.ID, &inputUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

// deleteUser remove a User
func (handler *UserHandler) deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	user, err := handler.UserService.User(cloudnativego.UserID(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No User found!"})
		return
	}

	err = handler.UserService.DeleteUser(cloudnativego.UserID(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
