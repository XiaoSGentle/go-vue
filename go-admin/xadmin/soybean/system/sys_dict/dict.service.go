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

type ISysDictService interface {
	GetDictList(c *gin.Context, param *SysDictListParam) (resp SysDictListResp, err error)
	AddDict(c *gin.Context, param *AddOrUpDateSysDictParam) (err error)
	UpdateDict(c *gin.Context, id int32, param *AddOrUpDateSysDictParam) (err error)
	DeleteDict(c *gin.Context, ids []int32) (err error)
}

func NewSysDictService(db *gorm.DB) ISysDictService {
	return &SysDictService{db: db, query: query.Use(db),
		dataFn: xgorm.InjectRouter[model.SysDictDatum](db),
		dictFn: xgorm.InjectRouter[model.SysDictType](db),
	}
}

type SysDictService struct {
	db     *gorm.DB
	query  *query.Query
	dataFn xgorm.IRouterFunctions[model.SysDictDatum]
	dictFn xgorm.IRouterFunctions[model.SysDictType]
}

func (s SysDictService) GetDictList(c *gin.Context, param *SysDictListParam) (resp SysDictListResp, err error) {
	sysDictQuery := query.Use(s.db).SysDictType.WithContext(c)
	resp.PageResult.PageParam = param.PageParam
	queryResultList, totalCount, err := sysDictQuery.FindByPage((param.Current-1)*param.Size, param.Size)
	if err != nil {
		return
	}
	resp.Total = totalCount
	for _, dictType := range queryResultList {
		resp.Records = append(resp.Records, SysDictList{
			BaseRecord: baseType.BaseRecord{
				ID:         dictType.ID,
				CreateBy:   dictType.CreateBy,
				CreateTime: dictType.CreateTime.String(),
				UpdateBy:   dictType.UpdateBy,
				UpdateTime: dictType.UpdateTime.String(),
				Status:     dictType.Status,
			},
			Name:        dictType.Name,
			Code:        dictType.Code,
			Description: dictType.Description,
		})
	}
	return
}

func (s SysDictService) AddDict(c *gin.Context, param *AddOrUpDateSysDictParam) (err error) {
	sysDictQuery := query.Use(s.db).SysDictType
	count, err := sysDictQuery.WithContext(c).Where(sysDictQuery.Code.Eq(param.Code)).Count()
	if count >= 1 {
		return xerror.NewErrCode(xerror.CURD_DATA_EXIST_ERROR)
	}
	err = sysDictQuery.WithContext(c).Create(&model.SysDictType{
		Name:        param.Name,
		Code:        param.Code,
		Description: param.Description,
		Status:      param.Status,
		Version:     0,
		CreateBy:    "",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		UpdateBy:    "",
	})
	return

}

func (s SysDictService) UpdateDict(c *gin.Context, id int32, param *AddOrUpDateSysDictParam) (err error) {
	sysDictQuery := query.Use(s.db).SysDictType
	count, err := sysDictQuery.WithContext(c).Where(sysDictQuery.ID.Eq(id)).Count()
	if count < 1 {
		return xerror.NewErrCode(xerror.CURD_DATA_NOT_EXIST_ERROR)
	}
	operateUserInfo := xtoken.GetBindCustomPayload(c)
	updates, err := sysDictQuery.WithContext(c).Where(sysDictQuery.ID.Eq(id)).Updates(model.SysDictType{
		Name:        param.Name,
		Code:        param.Code,
		Description: param.Description,
		Status:      param.Status,
		UpdateTime:  time.Now(),
		UpdateBy:    operateUserInfo.NickName,
		UpdateUID:   0,
	})
	if err != nil {
		return
	}
	if updates.RowsAffected < 1 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}

func (s SysDictService) DeleteDict(c *gin.Context, ids []int32) (err error) {
	menuQuery := s.query.SysDictType
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
