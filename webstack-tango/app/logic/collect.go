package logic

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"taogin/app/model"
	"taogin/app/types"
	"taogin/config/global"
	"time"
)

type CollectLogic struct {
}

func NewCollectLogic() *CollectLogic {
	return &CollectLogic{}
}

func (this *CollectLogic) GetList(req *types.CollectListReq) (resp *types.CollectListResp, err error) {
	var (
		total        uint64
		collectModel model.Collect
	)

	list := global.DB["default"].Model(collectModel)

	//获取总数
	if err := list.Count(&total).Error; err != nil {
		return nil, err
	}

	if len(req.Title) > 0 {
		list = list.Where("name", req.Title)
	}

	offset := (req.Page - 1) * req.PageSize
	list.Offset(offset).Limit(req.PageSize).Find(&collectModel)
	return &types.CollectListResp{
		Total:    total,
		Page:     0,
		PageSize: 0,
		//List:     collectModel,
	}, nil
}

func (this *CollectLogic) GetOne(id string) (resp *types.CollectBase) {

	collectModel := model.Collect{}
	collectBase := types.CollectBase{}

	global.DB["default"].First(&collectModel, id)
	collectBase = types.CollectBase{
		Id:    collectModel.Id,
		Title: collectModel.Title,
		Desc:  collectModel.Desc,
		CatId: collectModel.CatId,
		Cover: collectModel.Cover,
		Sort:  collectModel.Sort,
	}

	return &collectBase
}

func (this *CollectLogic) Add(req types.SaveCollectReq) error {
	var (
		collect model.Collect
	)
	//检查是否存在
	err := global.DB["default"].Table(collect.TableName()).Where("link = ?", req.Link).First(&collect).Error
	if err == gorm.ErrRecordNotFound {
		//无数据, 添加数据
		nowTime := time.Now()
		err := global.DB["default"].Table(collect.TableName()).Create(&model.Collect{
			Title:     req.Title,
			Desc:      req.Desc,
			CatId:     req.CatId,
			Cover:     req.Cover,
			Link:      req.Link,
			Sort:      req.Sort,
			CreatedAt: &nowTime,
		}).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return errors.New("操作失败")
	} else {
		return errors.New("已存在")
	}
	return nil
}

func (this *CollectLogic) Update(id string, req types.SaveCollectReq) error {
	var (
		collect      model.Collect
		collectOther model.Collect
	)
	//检查是否存在
	err := global.DB["default"].Table(collect.TableName()).Where("id = ?", id).First(&collect).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("不存在")
	} else if err != nil {
		return errors.New("操作失败")
	}
	//检查分类是否存在
	err = global.DB["default"].Table(collectOther.TableName()).Where("link = ?", req.Link).First(&collectOther).Error
	if err == gorm.ErrRecordNotFound {
		//更新
		err = global.DB["default"].Model(&collect).Where("id = ?", id).Updates(map[string]interface{}{
			"Title": req.Title,
			"Desc":  req.Desc,
			"CatId": req.CatId,
			"Cover": req.Cover,
			"Link":  req.Link,
			"Sort":  req.Sort,
		}).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return errors.New("操作失败")
	} else {
		return errors.New("已存在")
	}
	return nil
}

func (this *CollectLogic) Delete(id string) error {
	var (
		collect model.Collect
	)
	global.DB["default"].Where("id = ?", id).Delete(&collect)

	return nil
}

func (this *CollectLogic) Star(id string, userId uint64) error {
	var (
		userCollect model.UserCollect
		collect     model.Collect
	)

	//检查收录是否存在
	err := global.DB["default"].Table(collect.TableName()).Where("id = ?", id).First(&collect).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("该收录不存在")
	}
	if err != gorm.ErrRecordNotFound && err != nil {
		return errors.New("操作失败#4")
	}

	//查询是否已经收藏
	err = global.DB["default"].Where("collect_id = ?", id).Where("user_id = ?", userId).First(&userCollect).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return errors.New("操作失败#2")
	}
	//WebStack-Tango
	if err == gorm.ErrRecordNotFound {
		//转换成uint64
		collectId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return errors.New("操作失败#3")
		}
		//收藏
		err = global.DB["default"].Table(userCollect.TableName()).Create(&model.UserCollect{
			UserId:    userId,
			CollectId: collectId,
		}).Error
		if err != nil {
			return err
		}
	} else {
		//取消收藏
		global.DB["default"].Table(userCollect.TableName()).Where("id = ?", userCollect.Id).Delete(&userCollect)
	}

	return nil
}

func (this *CollectLogic) GetStar(userId uint64) (resp types.CollectStarResp) {
	var ()

	return resp
}
