package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/account/api/user"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	accountGroup := router.Group("/account/api/v1")

	user.RegisterRoutes(accountGroup)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
		c.Abort()
	})
	return router

}
