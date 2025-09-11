package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.RouterGroup) {
	userService := UserService{}
	userController := UserController{service: userService}

	route.GET("/users/me", userController.GetCurrentUser)
	route.POST("/users/onboard", userController.OnboardUser)
	route.GET("/users", userController.GetUsers)
	route.GET("/users/:id", userController.GetUser)
	route.POST("/users", userController.PostUser)
	route.DELETE("/users/:id", userController.DeleteUser)
	route.PATCH("/users/:id", userController.UpdateUser)

}
