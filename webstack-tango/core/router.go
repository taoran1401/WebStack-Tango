package core

import (
	"github.com/gin-gonic/gin"
	"taogin/app/middleware"
	default_router "taogin/router"
)

type Router struct {
}

//加载路由
func (this *Router) Routers() (router *gin.Engine) {
	engine := gin.Default()
	//全局中间件
	engine.Use(middleware.NewCorsMiddleware().Handle)
	//测试路由
	this.BaseRouter(engine)
	//路由
	this.InitRouter(engine)
	//websocket路由
	this.WsInitRouter()
	return engine
}

//加载路由
func (this *Router) InitRouter(router *gin.Engine) {
	default_router.InitRoute(router.Group(""))
}

//加载ws路由
func (this *Router) WsInitRouter() {
	default_router.WsInitRoute()
}

//测试路由
func (this *Router) BaseRouter(router *gin.Engine) {
	//test router
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})
}
