package logic

import (
	"errors"
	"github.com/jinzhu/gorm"
	"taogin/app/model"
	"taogin/app/types"
	"taogin/config/global"
	"time"
)

type RemindLogic struct {
}

func NewRemindLogic() *RemindLogic {
	return &RemindLogic{}
}

/*
func (this *RemindLogic) GetList(req *types.CollectListReq) (resp *types.CollectListResp, err error) {
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
}*/
/*
func (this *RemindLogic) GetOne(id string) (resp *types.CollectBase) {

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
}*/

func (this *RemindLogic) Add(req types.SaveRemindReq, userId uint64) error {
	var (
		remind model.Remind
	)
	//检查是否达到上线
	limit := this.getRemindLimit()
	var count int
	err := global.DB["default"].Table(remind.TableName()).Where("user_id = ?", userId).Count(&count).Error
	if count > limit {
		return errors.New("提醒任务已达上限")
	}

	if err != nil {
		return errors.New(err.Error())
	}

	if req.Type != 0 || req.Type != 1 {
		return errors.New("参数错误：类型")
	}

	//入库
	if err == gorm.ErrRecordNotFound {
		//无数据, 添加数据
		nowTime := time.Now()
		err := global.DB["default"].Table(remind.TableName()).Create(&model.Remind{
			UserId:    userId,
			Account:   req.Account,
			Content:   req.Content,
			Type:      req.Type,
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

/*
func (this *RemindLogic) Update(id string, req types.SaveCollectReq) error {
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
}*/

func (this *RemindLogic) Delete(id string) error {
	var (
		remind model.Remind
	)
	global.DB["default"].Where("id = ?", id).Delete(&remind)

	return nil
}

//获取提醒限制数量
func (this *RemindLogic) getRemindLimit() int {
	return 1
}
