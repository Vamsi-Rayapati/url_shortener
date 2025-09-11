package shorten

import "github.com/gin-gonic/gin"

func RegisterRoutes(route *gin.RouterGroup) {
	ShortenService := ShortenService{}
	ShortenController := ShortenController{service: ShortenService}

	route.PATCH("/users/:id", ShortenController.CreateShortURL)

}
