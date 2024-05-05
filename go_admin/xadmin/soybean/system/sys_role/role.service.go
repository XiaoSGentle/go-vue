package sys_role

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
	"xadmin/soybean/dao/model"
	baseType "xadmin/soybean/dao/model/base"
	"xadmin/soybean/dao/query"
	"xcore/common/xerror"
)

type ISysRoleService interface {
	SysRoleList(c *gin.Context, param *SysRoleListParam) (resp SysRoleListResp, err error)
	GetRolePermit(c *gin.Context, param string) (resp SysRolePermitResp, err error)
	AddRole(c *gin.Context, param *AddOrUpdateSysRoleParam) (err error)
	UpdateRole(c *gin.Context, id int32, param *AddOrUpdateSysRoleParam) (err error)
	DeleteRole(c *gin.Context, ids []int32) (err error)
	UpdateMenus(c *gin.Context, param UpdateRolePermitParam) (err error)
	UpdateApis(c *gin.Context, param UpdateRolePermitParam) (err error)
	UpdateHome(c *gin.Context, param UpdateRoleHomeParam) (err error)
}

func NewSysRoleService(db *gorm.DB) ISysRoleService {
	return &SysRoleService{db: db, query: query.Use(db)}
}

type SysRoleService struct {
	db    *gorm.DB
	query *query.Query
}

func (s SysRoleService) UpdateHome(c *gin.Context, param UpdateRoleHomeParam) (err error) {
	menuQuery := s.query.SysRole
	update, err := menuQuery.WithContext(c).Where(menuQuery.Code.Eq(param.RoleCode)).Update(menuQuery.Home, param.Home)
	if err != nil {
		return err
	}
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (s SysRoleService) UpdateMenus(c *gin.Context, param UpdateRolePermitParam) (err error) {
	menuQuery := s.query.SysRole
	update, err := menuQuery.WithContext(c).Where(menuQuery.Code.Eq(param.RoleCode)).Update(menuQuery.MenuIds, strings.Join(param.MenuIds, ","))
	if err != nil {
		return err
	}
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (s SysRoleService) UpdateApis(c *gin.Context, param UpdateRolePermitParam) (err error) {
	menuQuery := s.query.SysRole
	update, err := menuQuery.WithContext(c).Where(menuQuery.Code.Eq(param.RoleCode)).Update(menuQuery.APICodes, strings.Join(param.ApiCodes, ","))
	if err != nil {
		return err
	}
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (s SysRoleService) GetRolePermit(c *gin.Context, param string) (resp SysRolePermitResp, err error) {
	menuQuery := s.query.SysRole
	first, err := menuQuery.WithContext(c).Where(menuQuery.Code.Eq(param)).First()
	if err != nil {
		return SysRolePermitResp{}, err
	}
	resp = SysRolePermitResp{
		ApiCodes: strings.Split(first.APICodes, ","),
		MenuIds:  strings.Split(first.MenuIds, ","),
	}
	return
}

func (s SysRoleService) AddRole(c *gin.Context, param *AddOrUpdateSysRoleParam) (err error) {
	menuQuery := s.query.SysRole
	err = menuQuery.WithContext(c).Create(&model.SysRole{
		Name:          param.RoleName,
		Code:          param.RoleCode,
		ParentID:      0,
		Description:   param.RoleDesc,
		Status:        param.Status,
		MenuIds:       "",
		APICodes:      "",
		Version:       0,
		SoftDeleteTag: 0,
		UpdateTime:    time.Now(),
		UpdateUID:     0,
		CreateUID:     0,
		CreateBy:      "",
		CreateTime:    time.Now(),
		UpdateBy:      "",
	})
	return
}

func (s SysRoleService) UpdateRole(c *gin.Context, id int32, param *AddOrUpdateSysRoleParam) (err error) {
	menuQuery := s.query.SysRole
	updates, err := menuQuery.WithContext(c).Where(menuQuery.ID.Eq(id)).Updates(&model.SysRole{
		Name:        param.RoleName,
		Code:        param.RoleCode,
		Description: param.RoleDesc,
		Status:      param.Status,
		UpdateTime:  time.Now(),
		UpdateBy:    "",
	})
	if err != nil {
		return err
	}
	if updates.Error != nil {
		return updates.Error
	}
	return
}

func (s SysRoleService) DeleteRole(c *gin.Context, ids []int32) (err error) {
	menuQuery := s.query.SysRole
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

func (s SysRoleService) SysRoleList(c *gin.Context, param *SysRoleListParam) (sysRoleListResp SysRoleListResp, err error) {
	sysRoleQuery := query.Use(s.db).SysRole.WithContext(c)
	sysRoleListResp.PageResult.PageParam = param.PageParam
	roleListInSql, totalCount, err := sysRoleQuery.FindByPage((param.Current-1)*param.Size, param.Size)
	var userList []SysRoleList
	for _, role := range roleListInSql {
		userList = append(userList, SysRoleList{
			BaseRecord: baseType.BaseRecord{
				ID:         role.ID,
				CreateBy:   role.CreateBy,
				CreateTime: role.CreateTime.String(),
				UpdateBy:   role.UpdateBy,
				UpdateTime: role.UpdateTime.String(),
				Status:     role.Status,
			},
			RoleName: role.Name,
			RoleCode: role.Code,
			RoleDesc: role.Description,
			RoleHome: role.Home,
		})
	}
	sysRoleListResp.Records = userList
	sysRoleListResp.Total = totalCount
	return
}
