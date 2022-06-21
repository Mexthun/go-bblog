package router

import (
	"me-bblog/api"
	"me-bblog/views"
)

func (r *RouterStuct) loadLoginRouters() {
	r.GET("/login", views.HTML.Login)
	r.POST("/api/v1/login", api.API.Login)
}
