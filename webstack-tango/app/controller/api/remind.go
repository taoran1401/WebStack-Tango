package api

import (
	"github.com/gin-gonic/gin"
	"taogin/app/logic"
	"taogin/app/types"
	"taogin/config/atom"
	"taogin/core/response"
)

type RemindController struct {
}

func NewRemindController() *RemindController {
	return &RemindController{}
}

//5个提醒/每用户

func (this *RemindController) Index(ctx *gin.Context) {

}

func (this *RemindController) Show(ctx *gin.Context) {

}

func (this *RemindController) Store(ctx *gin.Context) {
	var (
		req types.SaveRemindReq
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}
	
	//参数验证
	if len(req.Account) == 0 || len(req.Content) == 0 {
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	//参数验证 - type
	typeMap := map[int64]bool{
		0: true,
		1: true,
	}
	if typeMap[req.Type] != true {
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	ctxUserInfo, _ := ctx.Get("user_info")
	userInfo, ok := ctxUserInfo.(*types.UserBase)
	if !ok {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, "操作失败#1")
		return
	}
	err := logic.NewRemindLogic().Add(req, userInfo.Id)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}
	response.Success(ctx, []string{})
}

func (this *RemindController) Update(ctx *gin.Context) {
	//
}

func (this *RemindController) Destroy(ctx *gin.Context) {
	//
}
