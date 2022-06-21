package api

import (
	"github.com/gin-gonic/gin"
	"me-bblog/common"
	"me-bblog/service"
	"me-bblog/utils"
)

func (*Api) Login(c *gin.Context) {
	data, getParamErr := common.GetRequestJsonParam(c.Request)
	if getParamErr != nil {
		panic(getParamErr)
	}
	username := data["username"].(string)
	passwd := data["passwd"].(string)
	passwd = utils.Md5Crypt(passwd, "mszlu")
	loginRes, err := service.Login(username, passwd)

	if err != nil {
		common.Error(c.Writer, err)
	} else {
		common.Success(c.Writer, loginRes)
	}

}
