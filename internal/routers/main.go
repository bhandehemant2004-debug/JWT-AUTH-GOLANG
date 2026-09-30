package routers

import "github.com/go-chi/chi"

type RouterInterface interface {
	Register(router *chi.Mux)
}
