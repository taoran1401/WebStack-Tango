package logic

import (
	"errors"
	"github.com/jinzhu/gorm"
	"taogin/app/model"
	"taogin/app/types"
	"taogin/config/global"
	"time"
)

type CollectCategoryLogic struct {
}

func NewCollectCategoryLogic() *CollectCategoryLogic {
	return &CollectCategoryLogic{}
}

func (this *CollectCategoryLogic) GetList(req *types.CollectCategoryListReq) (resp *types.CollectCategoryListResp, err error) {
	var (
		total                uint64
		collectCategoryModel model.CollectCategory
	)

	list := global.DB["default"].Model(collectCategoryModel)

	//获取总数
	if err := list.Count(&total).Error; err != nil {
		return nil, err
	}

	if len(req.Name) > 0 {
		list = list.Where("name", req.Name)
	}

	offset := (req.Page - 1) * req.PageSize
	list.Offset(offset).Limit(req.PageSize).Find(&collectCategoryModel)
	return &types.CollectCategoryListResp{
		Total:    total,
		Page:     0,
		PageSize: 0,
		//List:     collectCategoryModel,
	}, nil
}

func (this *CollectCategoryLogic) GetOne(id string) (resp *types.CollectCategoryBase) {

	collectCategoryModel := model.CollectCategory{}
	collectCategoryBase := types.CollectCategoryBase{}

	global.DB["default"].First(&collectCategoryModel, id)
	collectCategoryBase = types.CollectCategoryBase{
		Id:   collectCategoryModel.Id,
		Name: collectCategoryModel.Name,
		PId:  collectCategoryModel.PId,
		Sort: collectCategoryModel.Sort,
	}

	return &collectCategoryBase
}

func (this *CollectCategoryLogic) Add(req types.SaveCollectCategoryReq) error {
	var (
		collectCategory model.CollectCategory
	)
	//检查是否存在
	err := global.DB["default"].Table(collectCategory.TableName()).Where("name = ?", req.Name).First(&collectCategory).Error
	if err == gorm.ErrRecordNotFound {
		//无数据, 添加数据
		nowTime := time.Now()
		err := global.DB["default"].Table(collectCategory.TableName()).Create(&model.CollectCategory{
			Name:      req.Name,
			Sort:      req.Sort,
			PId:       0,
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

func (this *CollectCategoryLogic) Update(id string, req types.SaveCollectCategoryReq) error {
	var (
		collectCategory      model.CollectCategory
		collectCategoryOther model.CollectCategory
	)
	//检查是否存在
	err := global.DB["default"].Table(collectCategory.TableName()).Where("id = ?", id).First(&collectCategory).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("不存在")
	} else if err != nil {
		return errors.New("操作失败")
	}
	//检查分类是否存在
	err = global.DB["default"].Table(collectCategoryOther.TableName()).Where("name = ?", req.Name).First(&collectCategoryOther).Error
	if err == gorm.ErrRecordNotFound {
		//更新
		err = global.DB["default"].Model(&collectCategory).Where("id = ?", id).Updates(map[string]interface{}{
			"name": req.Name,
			"sort": req.Sort,
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

func (this *CollectCategoryLogic) Delete(id string) error {
	var (
		collectCategory model.CollectCategory
	)
	global.DB["default"].Where("id = ?", id).Delete(&collectCategory)

	return nil
}
