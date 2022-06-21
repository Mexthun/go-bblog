package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"me-bblog/common"
	"me-bblog/models"
	"me-bblog/service"
	"me-bblog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(c *gin.Context) {
	path := c.Request.URL.Path
	pidstr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		common.Error(c.Writer, err)
	}
	method := c.Request.Method
	switch method {
	case http.MethodDelete:
		err := service.DeletePostByPostId(pid)
		if err != nil {
			common.Error(c.Writer, err)
		}
		common.Success(c.Writer, "删除成功")
	case http.MethodGet:
		post, err := service.GetPost(pid)
		if err != nil {
			common.Error(c.Writer, err)
		}
		common.Success(c.Writer, post)

	}

}

func (*Api) SearchPost(c *gin.Context) {
	_ = c.Request.ParseForm()
	condition := c.Request.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(c.Writer, searchResp)
}
func (receiver Api) SaveAndUpdatePost(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(c.Writer, errors.New("登录已过期"))
	}
	uid := claim.Uid
	method := c.Request.Method

	params, errOfGetParam := common.GetRequestJsonParam(c.Request)
	if errOfGetParam != nil {
		common.Error(c.Writer, errors.New("服务器错误"))
	}
	switch method {
	case http.MethodPost:

		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		err := service.SavePost(post)
		if err != nil {
			common.Error(c.Writer, err)
		}
		common.Success(c.Writer, post)
	case http.MethodPut:
		// update

		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(c.Writer, post)
	}

}

func (*Api) Delete(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	_, _, errOfParseToken := utils.ParseToken(token)
	if errOfParseToken != nil {
		common.Error(c.Writer, errors.New("登录已过期"))
	}

	pidStr := c.Param("pid")
	pid, errPidStr2Int := strconv.Atoi(pidStr)
	if errPidStr2Int != nil {
		common.Error(c.Writer, errPidStr2Int)
	}
	err := service.DeletePostByPostId(pid)
	if err != nil {
		common.Error(c.Writer, err)
	}
	common.Success(c.Writer, "删除成功")
}
