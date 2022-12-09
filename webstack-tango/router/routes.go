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

	apiGroup := RouterGroup.Group("api")
	{
		//logout
		//login.POST("logout", api.NewLoginController().Logout)
		//index
		apiGroup.GET("/index", api.NewIndexController().Index).Use(middleware.NewJwtAuthMiddleware().Handle())
		//upload
		apiGroup.POST("uploads", api.NewUpload().Upload).Use(middleware.NewJwtAuthMiddleware().Handle())
		//user
		apiGroup.GET("users/:id", api.NewUserController().Show).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.PUT("users/:id", api.NewUserController().Update).Use(middleware.NewJwtAuthMiddleware().Handle())
		//collect category
		apiGroup.GET("collect/categorys", api.NewCollectCategory().Index).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.GET("collect/category/:id", api.NewCollectCategory().Show).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.POST("collect/category", api.NewCollectCategory().Store).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.PUT("collect/category/:id", api.NewCollectCategory().Update).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.DELETE("collect/category/:id", api.NewCollectCategory().Destroy).Use(middleware.NewJwtAuthMiddleware().Handle())
		//collect
		apiGroup.GET("collects", api.NewCollect().Index).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.GET("collect/:id", api.NewCollect().Show).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.POST("collect", api.NewCollect().Store).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.PUT("collect/:id", api.NewCollect().Update).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.DELETE("collect/:id", api.NewCollect().Destroy).Use(middleware.NewJwtAuthMiddleware().Handle())
		//action
		apiGroup.GET("collect/stars", api.NewCollect().GetStar).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.PUT("collect/star/:id", api.NewCollect().Star).Use(middleware.NewJwtAuthMiddleware().Handle())
		//反馈：tuxiaochao

		//short_link
		apiGroup.GET("reminds", api.NewRemindController().Index).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.GET("reminds/:id", api.NewRemindController().Show).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.POST("reminds", api.NewRemindController().Store).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.PUT("reminds/:id", api.NewRemindController().Update).Use(middleware.NewJwtAuthMiddleware().Handle())
		apiGroup.DELETE("reminds/:id", api.NewRemindController().Destroy).Use(middleware.NewJwtAuthMiddleware().Handle())
	}

	//websocket router
	/*ws := RouterGroup.Group("/ws")
	{
		ws.GET("/:token", ws2.WsHandler)
	}*/
}
