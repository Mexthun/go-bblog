package views

import (
	"github.com/gin-gonic/gin"
	"me-bblog/config"
	"net/http"
)

func (*HtmlApi) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", config.Cfg.Viewer)
}
