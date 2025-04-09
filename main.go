package main

import (
	"holamundo/core"

	userUsecase "holamundo/users/aplication/usecase"
	userRepo "holamundo/users/domain/repository"
	userControllers "holamundo/users/infraestructure/controllers"
	userRoutes "holamundo/users/infraestructure/routes"
	
	"github.com/gin-contrib/cors"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	core.InitializeApp()

	

	userRepo := &userRepo.UserRepositoryImpl{DB: core.GetDB()}
	createUserUC := &userUsecase.CreateUserUseCase{UserRepo: userRepo}
	getAllUsersUC := &userUsecase.GetAllUsersUseCase{UserRepo: userRepo}
	getUserUC := &userUsecase.GetUserUseCase{UserRepo: userRepo}
	updateUserUC := &userUsecase.UpdateUserUseCase{UserRepo: userRepo}
	deleteUserUC := &userUsecase.DeleteUserUseCase{UserRepo: userRepo}

	userCreateController := &userControllers.UserCreateController{CreateUserUC: createUserUC}
	userGetAllController := &userControllers.UserGetAllController{GetAllUsersUC: getAllUsersUC}
	userGetController := &userControllers.UserGetController{GetUserUC: getUserUC}
	userUpdateController := &userControllers.UserUpdateController{UpdateUserUC: updateUserUC}
	userDeleteController := &userControllers.UserDeleteController{DeleteUserUC: deleteUserUC}

	userRoutes.UserRoutes(app, userCreateController, userGetAllController, userUpdateController, userDeleteController, userGetController)

	log.Println("API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}