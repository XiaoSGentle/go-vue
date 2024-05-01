package sys_user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	baseType "xcore/dao/model/base"
	"xcore/dao/query"
)

type ISysUserService interface {
	SysUserList(c *gin.Context, param *SysUserListParam) (loginVo SysUserListResp, err error)
}

func NewSysUserService(db *gorm.DB) ISysUserService {
	return &SysUserService{db: db}
}

type SysUserService struct {
	db *gorm.DB
}

func (a SysUserService) SysUserList(c *gin.Context, param *SysUserListParam) (sysUserListResp SysUserListResp, err error) {
	sysUserQuery := query.Use(a.db).SysUser.WithContext(c)
	sysUserListResp.PageResult.PageParam = param.PageParam
	userListInSql, totalCount, err := sysUserQuery.FindByPage((param.Current-1)*param.Size, param.Size)
	var userList []SysUserList
	for _, user := range userListInSql {
		userList = append(userList, SysUserList{
			BaseRecord: baseType.BaseRecord{
				ID:         user.ID,
				CreateBy:   user.CreateBy,
				CreateTime: user.CreateTime.String(),
				UpdateBy:   user.UpdateBy,
				UpdateTime: user.UpdateTime.String(),
				Status:     user.UserStatus,
			},
			UserName:      user.Username,
			UserGender:    user.Gender,
			NickName:      user.Nickname,
			UserPhone:     user.Phone,
			UserEmail:     user.Email,
			UserRoles:     strings.Split(user.Roles, ","),
			LastOnLine:    user.LastOnlineTime.String(),
			LastCpWd:      user.LastCpwdTime.String(),
			NeedChangePwd: string(user.NeedChangePwd),
		})
	}
	sysUserListResp.Records = userList
	sysUserListResp.Total = totalCount
	return
}
