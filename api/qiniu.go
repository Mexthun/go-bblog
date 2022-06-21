package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"me-bblog/common"
	"me-bblog/config"
)

func (*Api) QiniuToken(c *gin.Context) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	//空间名称
	bucket := "methun"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(c.Writer, upToken)
}
