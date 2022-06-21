package router

import (
	"me-bblog/api"
	"me-bblog/views"
)

func (r *RouterStuct) post() {
	r.GET("/p/:pid", views.HTML.Detail)
	r.GET("/writing", views.HTML.Writing)

	r.POST("/api/v1/post", api.API.SaveAndUpdatePost)
	r.PUT("/api/v1/post", api.API.SaveAndUpdatePost)
	r.GET("/api/v1/post/:pid", api.API.GetPost)
	r.DELETE("/api/v1/post/:pid", api.API.Delete)
	r.GET("/api/v1/post/search", api.API.SearchPost)
}
