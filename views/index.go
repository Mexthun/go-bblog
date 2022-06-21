package views

import (
	"github.com/gin-gonic/gin"
	"me-bblog/service"
	"net/http"
	"strconv"
)

func (*HtmlApi) Index(c *gin.Context) {
	page := c.Query("page")
	var pageInt int
	var errPidStr2Int error
	if page != "" {
		pageInt, errPidStr2Int = strconv.Atoi(page)
		if errPidStr2Int != nil {
			c.HTML(http.StatusOK, "index.html", errPidStr2Int)
		}
	} else {
		pageInt = 1
	}

	homeRe, err := service.GetAllIndexInfo("", pageInt, 10)

	if err != nil {
		c.HTML(http.StatusOK, "index.html", err)

	}
	c.HTML(http.StatusOK, "index.html", homeRe)
}
