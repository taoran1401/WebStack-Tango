package api

import (
	"github.com/gin-gonic/gin"
	"taogin/app/logic"
	"taogin/app/types"
	"taogin/config/atom"
	"taogin/core/response"
)

type CollectCategory struct {
}

func NewCollectCategory() *CollectCategory {
	return &CollectCategory{}
}

func (this *CollectCategory) Show(ctx *gin.Context) {
	var (
		resp *types.CollectCategoryBase
	)
	id := ctx.Param("id")

	resp = logic.NewCollectCategoryLogic().GetOne(id)
	response.Success(ctx, resp)
}

func (this *CollectCategory) Store(ctx *gin.Context) {
	var (
		req types.SaveCollectCategoryReq
	)
	//解析json
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}

	if len(req.Name) == 0 {
		//参数错误
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	err := logic.NewCollectCategoryLogic().Add(req)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}
	response.Success(ctx, atom.SUCCESS)
}

func (this *CollectCategory) Update(ctx *gin.Context) {
	var (
		req types.SaveCollectCategoryReq
	)

	id := ctx.Param("id")

	//解析json
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, atom.ERROR_CODE_JSON, err.Error())
		return
	}

	if len(req.Name) == 0 {
		//参数错误
		response.Error(ctx, atom.ERROR_CODE_PARAM, atom.GetMsgByCode(atom.ERROR_CODE_PARAM))
		return
	}

	//更新
	err := logic.NewCollectCategoryLogic().Update(id, req)
	if err != nil {
		response.Error(ctx, atom.ERROR_CODE_EXCEPTION, err.Error())
		return
	}

	response.Success(ctx, []string{})
}

func (this *CollectCategory) Destroy(ctx *gin.Context) {

	id := ctx.Param("id")

	logic.NewCollectCategoryLogic().Delete(id)

	response.Success(ctx, []string{})
}
