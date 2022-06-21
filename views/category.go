package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"me-bblog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Category(c *gin.Context) {

	pageStr := c.Request.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryName := c.Param("categoryName")

	var cId int
	var err error
	if categoryName != "" {

	} else {
		//http://localhost:8080/c/1  1参数 分类的id
		path := c.Request.URL.Path
		cIdStr := strings.TrimPrefix(path, "/c/")
		cId, err = strconv.Atoi(cIdStr)
		if err != nil {
			c.HTML(http.StatusOK, "category.html", errors.New("不识别此请求路径"))
			return
		}
		if err := c.Request.ParseForm(); err != nil {
			c.HTML(http.StatusOK, "category.html", errors.New("系统错误，请联系管理员!!"))
			return
		}
	}

	categoryResponse, errOfGetCategoryResponse := service.GetPostsByCategoryIdOrCategoryName(categoryName, cId, page, pageSize)
	if errOfGetCategoryResponse != nil {
		c.HTML(http.StatusOK, "category.html", err)
		return
	}
	c.HTML(http.StatusOK, "category.html", categoryResponse)

}
