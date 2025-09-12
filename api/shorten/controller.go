package shorten

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/url_shortener/pkg/validator"
)

type ShortenController struct {
	service ShortenService
}

func (sc *ShortenController) CreateShortURL(c *gin.Context) {
	var request ShortenRequest
	err := validator.ValidateBody(c, &request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := sc.service.CreateShortURL(request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}
