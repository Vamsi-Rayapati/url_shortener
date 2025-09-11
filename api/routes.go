package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/url_shortener/api/shorten"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	group := router.Group("/api/v1")

	shorten.RegisterRoutes(group)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
		c.Abort()
	})
	return router

}
