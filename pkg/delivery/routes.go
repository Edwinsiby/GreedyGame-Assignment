package delivery

import (
	"greedy/pkg/domain"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handlers HandlerMethods
}

type RoutesMethods interface {
	SetKeyValueRoutes(*gin.Engine)
	SetQueRoutes(*gin.Engine)
}

func NewRoutes(config *domain.Config, handlers HandlerMethods) RoutesMethods {

	return Routes{
		handlers: handlers,
	}

}

func (r Routes) SetKeyValueRoutes(router *gin.Engine) {
	router.POST("key/set", r.handlers.Set)
	router.GET("key/get", r.handlers.Get)
}

func (r Routes) SetQueRoutes(router *gin.Engine) {
	router.POST("que/push", r.handlers.QPush)
	router.GET("que/pop", r.handlers.QPop)
	router.GET("que/bpop", r.handlers.BQPop)
}
