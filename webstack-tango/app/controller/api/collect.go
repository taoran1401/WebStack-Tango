package api

import (
	"github.com/gin-gonic/gin"
	"taogin/app/logic"
	"taogin/app/types"
	"taogin/config/atom"
	"taogin/core/response"
)

type Collect struct {
}

func NewCollect() *Collect {
	return &Collect{}
}

func (this *Collect) Index(ctx *gin.Context) {
	var (
		req types.CollectListReq
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}

	logic.NewCollectLogic().GetList(&req)
}

func (this *Collect) Show(ctx *gin.Context) {
	var (
		resp *types.CollectBase
	)
	id := ctx.Param("id")

	resp = logic.NewCollectLogic().GetOne(id)
	response.Success(ctx, resp)
}

func (this *Collect) Store(ctx *gin.Context) {
	var (
		req types.SaveCollectReq
	)
	//解析json
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}

	if len(req.Link) == 0 || len(req.Title) == 0 {
		//参数错误
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	err := logic.NewCollectLogic().Add(req)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}
	response.Success(ctx, atom.SUCCESS)
}

func (this *Collect) Update(ctx *gin.Context) {
	var (
		req types.SaveCollectReq
	)

	id := ctx.Param("id")

	//解析json
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}

	if len(req.Link) == 0 || len(req.Title) == 0 {
		//参数错误
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	//更新
	err := logic.NewCollectLogic().Update(id, req)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}

	response.Success(ctx, []string{})
}

func (this *Collect) Destroy(ctx *gin.Context) {

	id := ctx.Param("id")

	logic.NewCollectLogic().Delete(id)

	response.Success(ctx, []string{})
}

func (this *Collect) Star(ctx *gin.Context) {
	id := ctx.Param("id")
	ctxUserInfo, _ := ctx.Get("user_info")
	userInfo, ok := ctxUserInfo.(*types.UserBase)
	if !ok {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, "操作失败#1")
		return
	}
	err := logic.NewCollectLogic().Star(id, userInfo.Id)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}
	response.Success(ctx, []string{})
}

func (this *Collect) GetStar(ctx *gin.Context) {
	var (
	//req types.CollectGetStarReq
	)
	ctxUserInfo, _ := ctx.Get("user_info")
	userInfo, ok := ctxUserInfo.(types.UserBase)
	if !ok {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, "操作失败")
		return
	}
	resp := logic.NewCollectLogic().GetStar(userInfo.Id)

	response.Success(ctx, resp)
}
