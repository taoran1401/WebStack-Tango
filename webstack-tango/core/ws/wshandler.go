package ws

import (
	"github.com/gin-gonic/gin"
)

func WsHandler(ctx *gin.Context) {
	WsServer(ctx, ctx.Writer, ctx.Request)
}
