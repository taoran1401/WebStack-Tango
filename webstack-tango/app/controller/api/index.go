package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"taogin/app/model"
	"taogin/app/types"
	"taogin/config/global"
	"taogin/core/response"
	"time"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (this *IndexController) Index(ctx *gin.Context) {

	var (
		collectCategoryModel     model.CollectCategory
		collect                  model.Collect
		collectCategoryModelList []model.CollectCategory
		collectModelList         []model.Collect
		list                     map[string]interface{}
	)

	homePageKey := "home-page"
	homePageData := global.CACHE.Get(homePageKey)
	if len(homePageData) != 0 {
		//有缓存，直接使用缓存
		_ = json.Unmarshal([]byte(homePageData), &collectCategoryModelList)
	} else {
		//无缓存
		global.DB["default"].Table(collectCategoryModel.TableName()).Find(&collectCategoryModelList)
		for i, v := range collectCategoryModelList {
			global.DB["default"].Table(collect.TableName()).Where("cat_id = ?", v.Id).Find(&collectModelList)
			collectCategoryModelList[i].Collect = collectModelList
		}

		//添加缓存(失效周期1天)
		if homePageJson, err := json.Marshal(collectCategoryModelList); err == nil {
			_ = global.CACHE.Set(homePageKey, homePageJson, time.Second*3600*24)
		}
	}
	ctxUserInfo, _ := ctx.Get("user_info")
	userInfo, ok := ctxUserInfo.(types.UserBase)
	if !ok {
		userInfo = types.UserBase{}
	}
	//build data
	list = map[string]interface{}{
		"collectCategoryModelList": collectCategoryModelList,
		"userInfo":                 userInfo,
	}
	response.Success(ctx, list)
	/*ctx.HTML(http.StatusOK, "index.html", gin.H{
		"List": list,
	})*/
}

func (this *IndexController) About(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about.html", gin.H{})
}

func (this *IndexController) NotFund(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.html", gin.H{})
}
