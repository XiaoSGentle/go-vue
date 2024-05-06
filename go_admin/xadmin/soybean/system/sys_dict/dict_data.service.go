package sys_dict

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"xadmin/soybean/dao/model"
	baseType "xadmin/soybean/dao/model/base"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
	xtoken "xcore/common/xtoken/jwt"
)

type ISysDictDataService interface {
	GetDictDataList(c *gin.Context, param *SysDictDataListParam) (resp SysDictDataListResp, err error)
	AddDictData(c *gin.Context, param *AddOrUpDateSysDictDataParam) (err error)
	UpdateDictData(c *gin.Context, id int32, param *AddOrUpDateSysDictDataParam) (err error)
	DeleteDictData(c *gin.Context, ids []int32) (err error)
}

func NewSysDictDataService(db *gorm.DB) ISysDictDataService {
	return &SysDictDataService{db: db, query: query.Use(db)}
}

type SysDictDataService struct {
	db    *gorm.DB
	query *query.Query
}

func (s SysDictDataService) GetDictDataList(c *gin.Context, param *SysDictDataListParam) (resp SysDictDataListResp, err error) {
	dictDataQuery := s.query.SysDictDatum
	resp.PageResult.PageParam = param.PageParam
	queryResultList, totalCount, err := dictDataQuery.WithContext(c).
		Where(dictDataQuery.TypeCode.Eq(param.Code)).
		FindByPage((param.Current-1)*param.Size, param.Size)
	resp.Total = totalCount
	if err != nil {
		return SysDictDataListResp{}, err
	}
	for _, datum := range queryResultList {
		resp.Records = append(resp.Records, SysDictDataList{
			BaseRecord: baseType.BaseRecord{
				ID:         datum.ID,
				CreateBy:   datum.CreateBy,
				CreateTime: datum.CreateTime.String(),
				UpdateBy:   datum.UpdateBy,
				UpdateTime: datum.UpdateTime.String(),
				Status:     datum.Status,
			},
			Label: datum.Label,
			Value: datum.Value,
			Sort:  datum.Sort,
		})
	}
	return
}

func (s SysDictDataService) AddDictData(c *gin.Context, param *AddOrUpDateSysDictDataParam) (err error) {
	sysDictQuery := query.Use(s.db).SysDictDatum
	payload := xtoken.GetBindCustomPayload(c)

	err = sysDictQuery.WithContext(c).Create(&model.SysDictDatum{
		Label:      param.Label,
		Value:      param.Value,
		Sort:       param.Sort,
		TypeCode:   param.TypeCode,
		Status:     param.Status,
		CreateUID:  payload.Uid,
		CreateBy:   payload.NickName,
		CreateTime: time.Now(),
	})
	return
}

func (s SysDictDataService) UpdateDictData(c *gin.Context, id int32, param *AddOrUpDateSysDictDataParam) (err error) {
	sysDictQuery := query.Use(s.db).SysDictDatum
	count, err := sysDictQuery.WithContext(c).Where(sysDictQuery.ID.Eq(id)).Count()
	if count < 1 {
		return xerror.NewErrCode(xerror.CURD_DATA_NOT_EXIST_ERROR)
	}
	operateUserInfo := xtoken.GetBindCustomPayload(c)
	updates, err := sysDictQuery.WithContext(c).Where(sysDictQuery.ID.Eq(id)).Updates(model.SysDictDatum{
		Label:      param.Label,
		Value:      param.Value,
		Sort:       param.Sort,
		TypeCode:   param.TypeCode,
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
	menuQuery := s.query.SysDictDatum
	info, err := menuQuery.WithContext(c).Where(menuQuery.ID.In(ids...)).Delete()
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
