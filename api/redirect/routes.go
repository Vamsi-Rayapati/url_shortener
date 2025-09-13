package redirect

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	rs := RedirectService{}
	rc := RedirectController{
		service: rs,
	}
	router.GET("/:id", rc.Redirect)

}
