package views

import (
	"github.com/gin-gonic/gin"
	"me-bblog/service"
	"net/http"
)

func (*HtmlApi) Pigeonhole(c *gin.Context) {

	pigeonholeRes := service.FindPostPigeonhole()

	c.HTML(http.StatusOK, "pigeonhole.html", pigeonholeRes)
}
