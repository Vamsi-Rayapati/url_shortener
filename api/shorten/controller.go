package shorten

import "github.com/gin-gonic/gin"

type ShortenController struct {
	service ShortenService
}

func (sc *ShortenController) CreateShortURL(c *gin.Context) {
	// CreateShortURL logic here
}
