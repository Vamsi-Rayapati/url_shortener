package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	service RedirectService
}

func (rc *RedirectController) Redirect(c *gin.Context) {

	id := c.Param("id")
	url, err := rc.service.GetActualUrl(id)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Redirect(http.StatusPermanentRedirect, *url)
}
