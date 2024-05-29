package sys_department

import (
	"gorm.io/gorm"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xgorm"
)

type ISysDepartService interface {
	FindByName()
}

type SysDepartService struct {
	xgorm.IServiceFunctions[model.SysDepartment]
	query *query.Query
	db    *gorm.DB
}

func (s SysDepartService) FindByName() {
	//TODO implement me
	panic("implement me")
}

func NewSysDepartService(db *gorm.DB) ISysDepartService {
	return &SysDepartService{
		IServiceFunctions: xgorm.InjectService[model.SysDepartment](db),
		query:             query.Use(db),
		db:                db,
	}
}
