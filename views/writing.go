package views

import (
	"github.com/gin-gonic/gin"
	"me-bblog/service"
	"net/http"
)

func (*HtmlApi) Writing(c *gin.Context) {
	//idStr := c.Request.PostForm.Get("id")
	idStr := c.Query("id")
	if idStr != "" {

	}
	res, err := service.Writing()
	if err != nil {
		c.HTML(http.StatusOK, "writing.html", err)
	}
	c.HTML(http.StatusOK, "writing.html", res)
}
