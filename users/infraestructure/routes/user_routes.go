package routes

import (
	
	"holamundo/users/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func UserRoutes(
	router *gin.Engine,
	createController *controllers.UserCreateController,
	getAllController *controllers.UserGetAllController,
	updateController *controllers.UserUpdateController,
	deleteController *controllers.UserDeleteController,
	getController *controllers.UserGetController,

) {
	group := router.Group("/users")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)  
		group.GET("/:id", getController.GetUserByID)   
		group.PUT("/:id", updateController.Update)    
		group.DELETE("/:id", deleteController.Delete )   
    
	}
}