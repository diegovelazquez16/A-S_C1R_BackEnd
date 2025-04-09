package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/users/aplication/usecase"
)

type UserGetController struct {
	GetUserUC *usecase.GetUserUseCase
}

func (c *UserGetController) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número válido."})
		return
	}

	user, err := c.GetUserUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	// Devolver la respuesta
	ctx.JSON(http.StatusOK, user)
}
