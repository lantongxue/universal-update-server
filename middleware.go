package main

import "github.com/gin-gonic/gin"

type Middleware struct {
	engine *gin.Engine
}

func NewMiddleware(r *gin.Engine) *Middleware {
	return &Middleware{
		engine: r,
	}
}

func (m *Middleware) Register(middlewares map[string]gin.HandlerFunc) {
	for _, middleware := range middlewares {
		m.engine.Use(middleware)
	}
}
