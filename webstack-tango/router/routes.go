package router

import (
	"github.com/gin-gonic/gin"
	"taogin/app/controller/api"
	"taogin/app/middleware"
)

func InitRoute(RouterGroup *gin.RouterGroup) {

	/*webGroup := RouterGroup.Group("/") //.Use(middleware.NewJwtAuthMiddleware().Handle())
	{
		//首页
		webGroup.GET("/", web.NewIndexController().Index)
	}*/

	//设置路由组、中间件
	login := RouterGroup.Group("api")
	{
		login.POST("send/phone/code", api.NewLoginController().SendPhoneCode) //发送验证码
		login.POST("login", api.NewLoginController().Login)                   //登录
		login.POST("register", api.NewLoginController().Register)             //注册
	}

	apiGroup := RouterGroup.Group("api").Use(middleware.NewJwtAuthMiddleware().Handle())
	{
		//logout
		//login.POST("logout", api.NewLoginController().Logout)
		//index
		apiGroup.GET("/index", api.NewIndexController().Index)
		//upload
		apiGroup.POST("uploads", api.NewUpload().Upload)
		//user
		apiGroup.GET("users/:id", api.NewUserController().Show)
		apiGroup.PUT("users/:id", api.NewUserController().Update)
		//collect category
		apiGroup.GET("collect/categorys/:id", api.NewCollectCategory().Show)
		apiGroup.POST("collect/categorys", api.NewCollectCategory().Store)
		apiGroup.PUT("collect/categorys/:id", api.NewCollectCategory().Update)
		apiGroup.DELETE("collect/categorys/:id", api.NewCollectCategory().Destroy)
		//collect
		apiGroup.GET("collects", api.NewCollect().Index)
		apiGroup.GET("collect/:id", api.NewCollect().Show)
		apiGroup.POST("collect", api.NewCollect().Store)
		apiGroup.PUT("collect/:id", api.NewCollect().Update)
		apiGroup.DELETE("collect/:id", api.NewCollect().Destroy)
		//action
		apiGroup.GET("collect/stars", api.NewCollect().GetStar)
		apiGroup.PUT("collect/star/:id", api.NewCollect().Star)
		//反馈：tuxiaochao
	}

	//websocket router
	/*ws := RouterGroup.Group("/ws")
	{
		ws.GET("/:token", ws2.WsHandler)
	}*/
}
