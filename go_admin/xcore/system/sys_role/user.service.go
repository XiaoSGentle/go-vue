package sys_role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	baseType "xcore/dao/model/base"
	"xcore/dao/query"
)

type ISysRoleService interface {
	SysRoleList(c *gin.Context, param *SysRoleListParam) (loginVo SysRoleListResp, err error)
}

func NewSysRoleService(db *gorm.DB) ISysRoleService {
	return &SysRoleService{db: db}
}

type SysRoleService struct {
	db *gorm.DB
}

func (a SysRoleService) SysRoleList(c *gin.Context, param *SysRoleListParam) (sysRoleListResp SysRoleListResp, err error) {
	sysRoleQuery := query.Use(a.db).SysRole.WithContext(c)
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
				Status:     fmt.Sprintf("%d", role.Status),
			},
			RoleName: role.Name,
			RoleCode: role.Name,
			RoleDesc: role.Description,
		})
	}
	sysRoleListResp.Records = userList
	sysRoleListResp.Total = totalCount
	return
}
