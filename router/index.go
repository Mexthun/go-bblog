package router

import (
	"me-bblog/views"
)

func (r *RouterStuct) index() {
	r.GET("/", views.HTML.Index)
}
