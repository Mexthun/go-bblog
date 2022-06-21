package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"me-bblog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Detail(c *gin.Context) {
	path := c.Request.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		c.HTML(http.StatusOK, "detail.html", errors.New("不识别此请求路径"))
	}
	postRes, err := service.GetPostDetail(pid)
	c.HTML(http.StatusOK, "detail.html", postRes)

}
