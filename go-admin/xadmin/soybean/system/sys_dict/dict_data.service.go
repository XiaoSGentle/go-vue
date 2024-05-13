package sys_dict

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
	"xcore/common/xgorm"
	"xcore/common/xtoken"
	baseType "xcore/common/xtype/xbase"
)

type ISysDictDataService interface {
	GetDictDataList(c *gin.Context, param *SysDictDataListParam) (resp SysDictDataListResp, err error)
	AddDictData(c *gin.Context, param *AddOrUpDateSysDictDataParam) (err error)
	UpdateDictData(c *gin.Context, id int32, param *AddOrUpDateSysDictDataParam) (err error)
	DeleteDictData(c *gin.Context, ids []int32) (err error)
}

func NewSysDictDataService(db *gorm.DB) ISysDictDataService {
	return &SysDictDataService{db: db, query: query.Use(db), serviceFun: xgorm.InjectService[model.SysDictDatum](db)}
}

type SysDictDataService struct {
	db         *gorm.DB
	query      *query.Query
	serviceFun xgorm.IServiceFunctions[model.SysDictDatum]
}

func (s SysDictDataService) GetDictDataList(c *gin.Context, param *SysDictDataListParam) (resp SysDictDataListResp, err error) {
	var result SysDictDataListResp = SysDictDataListResp{
		Records: []SysDictDataList{},
	}
	dictDataQuery := s.query.SysDictDatum
	resp.PageResult.PageParam = param.PageParam
	queryResultList, totalCount, err := dictDataQuery.WithContext(c).
		Where(dictDataQuery.TypeCode.Eq(param.Code)).
		Where(dictDataQuery.DeleteTag.Eq(0)).
		FindByPage((param.Current-1)*param.Size, param.Size)
	resp.Total = totalCount
	if err != nil {
		return result, err
	}
	for _, datum := range queryResultList {
		result.Records = append(result.Records, SysDictDataList{
			BaseRecord: baseType.BaseRecord{
				ID:         datum.ID,
				CreateBy:   datum.CreateBy,
				CreateTime: datum.CreateTime.String(),
				UpdateBy:   datum.UpdateBy,
				UpdateTime: datum.UpdateTime.String(),
				Status:     datum.Status,
			},
			Label:   datum.Label,
			Value:   datum.Value,
			EnLabel: datum.EnLabel,
			Code:    param.Code,
			Sort:    datum.Sort,
		})
	}
	return result, nil
}

func (s SysDictDataService) AddDictData(c *gin.Context, param *AddOrUpDateSysDictDataParam) (err error) {
	dictDataQuery := query.Use(s.db).SysDictDatum
	payload := xtoken.GetBindCustomPayload(c)

	err = dictDataQuery.WithContext(c).Create(&model.SysDictDatum{
		Label:      param.Label,
		Value:      param.Value,
		Sort:       param.Sort,
		EnLabel:    param.EnLabel,
		TypeCode:   param.Code,
		Status:     param.Status,
		CreateUID:  payload.Uid,
		CreateBy:   payload.NickName,
		CreateTime: time.Now(),
	})
	return
}

func (s SysDictDataService) UpdateDictData(c *gin.Context, id int32, param *AddOrUpDateSysDictDataParam) (err error) {
	dictDataQuery := query.Use(s.db).SysDictDatum
	count, err := dictDataQuery.WithContext(c).
		Where(dictDataQuery.ID.Eq(id)).
		Where(dictDataQuery.DeleteTag.Eq(0)).
		Count()
	if count < 1 {
		return xerror.NewErrCode(xerror.CURD_DATA_NOT_EXIST_ERROR)
	}
	operateUserInfo := xtoken.GetBindCustomPayload(c)
	updates, err := dictDataQuery.WithContext(c).
		Where(dictDataQuery.ID.Eq(id)).
		Where(dictDataQuery.DeleteTag.Eq(0)).
		Updates(model.SysDictDatum{
			Label:      param.Label,
			Value:      param.Value,
			EnLabel:    param.EnLabel,
			Sort:       param.Sort,
			TypeCode:   param.Code,
			Status:     param.Status,
			UpdateTime: time.Now(),
			UpdateUID:  operateUserInfo.Uid,
			UpdateBy:   operateUserInfo.NickName,
		})
	if err != nil {
		return
	}
	if updates.RowsAffected < 1 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}

func (s SysDictDataService) DeleteDictData(c *gin.Context, ids []int32) (err error) {
	dictDataQuery := s.query.SysDictDatum
	operatorInfo := s.serviceFun.GetOperatorInfo(c)
	info, err := dictDataQuery.WithContext(c).
		Where(dictDataQuery.ID.In(ids...)).
		Updates(&model.SysDictDatum{
			DeleteTag:  1,
			UpdateBy:   operatorInfo.NickName,
			UpdateUID:  operatorInfo.Uid,
			UpdateTime: time.Now(),
		})
	if err != nil {
		return err
	}
	if info.Error != nil {
		return info.Error
	}
	if info.RowsAffected < 1 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}
