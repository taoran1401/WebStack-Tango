package response

import (
	"github.com/gin-gonic/gin"
	"taogin/config/atom"
)

func Success(ctx *gin.Context, res interface{}) {
	ctx.JSON(200, gin.H{
		"code": atom.SUCCESS,
		"msg":  atom.GetMsgByCode(atom.SUCCESS),
		"data": res,
	})
	ctx.Abort()
}

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
	ctx.Abort()
}

func HtmlSuccess(ctx *gin.Context, tmpl string, res interface{}) {
	ctx.HTML(200, tmpl, gin.H{
		"data": res,
	})
	ctx.Abort()
}
