package sys_cron

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xgorm"
	"xcore/common/xtype/xbase"
)

type ISysCronService interface {
	FindCronList(c *gin.Context, param *xbase.PageParam) (resp SysCronListResp, err error)
	UpdateCron(c *gin.Context, cron *UpdateCronReq) (err error)
}

type SysCronService struct {
	serviceFn xgorm.IServiceFunctions[model.SysCron]
	query     *query.Query
	db        *gorm.DB
}

func (s SysCronService) FindCronList(c *gin.Context, param *xbase.PageParam) (resp SysCronListResp, err error) {
	cronQuery := query.Use(s.db).SysCron
	resp.PageResult.PageParam = *param
	var respList []CronResp
	userListInSql, totalCount, err := cronQuery.WithContext(c).
		FindByPage((param.Current-1)*param.Size, param.Size)
	for _, cron := range userListInSql {
		respList = append(respList, CronResp{
			Key:         cron.Key,
			Description: cron.Description,
			Schedule:    cron.Schedule,
			Status:      cron.Status,
			Arguments:   strings.Split(cron.Arguments, ","),
		})
	}
	resp.Records = respList
	resp.Total = totalCount
	return
}

func (s SysCronService) UpdateCron(c *gin.Context, cron *UpdateCronReq) (err error) {
	cronQuery := query.Use(s.db).SysCron

	updates, err := cronQuery.WithContext(c).Where(cronQuery.Key.Eq(cron.Key)).Updates(model.SysCron{
		Key:         cron.Key,
		Description: cron.Description,
		Schedule:    cron.Schedule,
		Status:      cron.Status,
		Arguments:   strings.Join(cron.Arguments, ","),
	})
	if err != nil {
		return err
	}
	if updates.Error != nil {
		return updates.Error
	}
	return
}

func NewSysDepartService(db *gorm.DB) ISysCronService {
	return &SysCronService{
		db:        db,
		serviceFn: xgorm.InjectService[model.SysCron](db),
		query:     query.Use(db),
	}
}
