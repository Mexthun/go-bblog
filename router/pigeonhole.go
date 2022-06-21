package router

import "me-bblog/views"

func (r *RouterStuct) Pigeonhole() {

	r.GET("/pigeonhole", views.HTML.Pigeonhole)
}
