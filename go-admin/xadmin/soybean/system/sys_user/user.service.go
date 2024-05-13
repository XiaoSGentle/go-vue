package sys_user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
	"xadmin/soybean/dao/model"
	"xadmin/soybean/dao/query"
	"xcore/common/xencrypt"
	"xcore/common/xerror"
	"xcore/common/xgorm"
	baseType "xcore/common/xtype/xbase"
)

type ISysUserService interface {
	SysUserList(c *gin.Context, param *SysUserListParam) (loginVo SysUserListResp, err error)
	AddUser(c *gin.Context, param *AddOrUpdateSysUserParam) (err error)
	UpdateUser(c *gin.Context, id int32, param *AddOrUpdateSysUserParam) (err error)
	DeleteUser(c *gin.Context, ids []int32) (err error)
}

func NewSysUserService(db *gorm.DB) ISysUserService {
	return &SysUserService{db: db, query: query.Use(db), sysUserInjectFunction: xgorm.InjectRouter[model.SysUser](db)}
}

type SysUserService struct {
	db                    *gorm.DB
	query                 *query.Query
	sysUserInjectFunction xgorm.IRouterFunctions[model.SysUser]
}

func (s SysUserService) AddUser(c *gin.Context, param *AddOrUpdateSysUserParam) (err error) {
	menuQuery := s.query.SysUser
	err = menuQuery.WithContext(c).Create(&model.SysUser{
		Nickname:       param.NickName,
		Username:       param.UserName,
		Roles:          strings.Join(param.UserRoles, ","),
		Password:       xencrypt.Base64Md5(param.UserName),
		LoginAttempts:  time.Now().Format(time.DateOnly) + "|0",
		Phone:          param.UserPhone,
		Avatar:         "",
		Email:          param.UserEmail,
		NeedChangePwd:  1,
		Gender:         param.UserGender,
		LastOnlineTime: time.Now(),
		LastCpwdTime:   time.Now(),
		UserStatus:     param.Status,
		Version:        0,
		SoftDeleteTag:  0,
		UpdateTime:     time.Now(),
		UpdateUID:      0,
		CreateUID:      0,
		CreateBy:       "",
		CreateTime:     time.Now(),
		UpdateBy:       "",
	})
	return
}

func (s SysUserService) UpdateUser(c *gin.Context, id int32, param *AddOrUpdateSysUserParam) (err error) {
	menuQuery := s.query.SysUser
	updates, err := menuQuery.WithContext(c).Where(menuQuery.ID.Eq(id)).Updates(&model.SysUser{
		Nickname:       param.UserName,
		Username:       param.NickName,
		Roles:          strings.Join(param.UserRoles, ","),
		Password:       xencrypt.Base64Md5(param.UserName),
		LoginAttempts:  time.Now().Format(time.DateOnly) + "|0",
		Phone:          param.UserPhone,
		Avatar:         "",
		Email:          param.UserEmail,
		NeedChangePwd:  1,
		Gender:         param.UserGender,
		LastOnlineTime: time.Time{},
		LastCpwdTime:   time.Time{},
		UserStatus:     param.Status,
		Version:        0,
		SoftDeleteTag:  0,
		UpdateTime:     time.Now(),
		UpdateUID:      0,
		CreateUID:      0,
		CreateBy:       "",
		CreateTime:     time.Now(),
		UpdateBy:       "",
	})
	if err != nil {
		return err
	}
	if updates.Error != nil {
		return updates.Error
	}
	return
}

func (s SysUserService) DeleteUser(c *gin.Context, ids []int32) (err error) {
	menuQuery := s.query.SysUser
	info, err := menuQuery.WithContext(c).Where(menuQuery.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.Error != nil {
		return info.Error
	}
	if info.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}

func (s SysUserService) SysUserList(c *gin.Context, param *SysUserListParam) (sysUserListResp SysUserListResp, err error) {
	sysUserQuery := query.Use(s.db).SysUser.WithContext(c)
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
