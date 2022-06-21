package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"me-bblog/common"
)

type RouterStuct struct {
	*gin.Engine
}

var router RouterStuct

func Router() {
	router.Engine = gin.Default()

	router.SetFuncMap(template.FuncMap{"isODD": common.IsODD, "getNextName": common.GetNextName, "date": common.Date, "dateDay": common.DateDay})
	router.Static("/resource", "./public/resource")
	router.LoadHTMLGlob("template/**/*")
	router.loadLoginRouters()
	router.qiNiuRouters()
	router.Category()
	router.index()
	router.post()
	router.Pigeonhole()

	router.Run(":8080")
}
