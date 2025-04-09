package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/users/aplication/usecase"
	"holamundo/users/domain/models"
)
type UserUpdateController struct {
	UpdateUserUC *usecase.UpdateUserUseCase
}

func (c *UserUpdateController) Update(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user.ID = uint(id) 
	err = c.UpdateUserUC.Execute(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}