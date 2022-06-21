package router

import (
	"me-bblog/views"
)

func (r *RouterStuct) Category() {
	r.GET("/c/:cid", views.HTML.Category)
	r.GET("/category/:categoryName", views.HTML.Category)
}
