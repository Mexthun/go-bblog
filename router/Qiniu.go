package router

import (
	"me-bblog/api"
)

func (r *RouterStuct) qiNiuRouters() {
	r.GET("/api/v1/qiniu/token", api.API.QiniuToken)
}
