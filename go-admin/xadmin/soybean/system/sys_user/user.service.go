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
	return &SysUserService{db: db, query: query.Use(db), serviceFun: xgorm.InjectService[model.SysUser](db)}
}

type SysUserService struct {
	db         *gorm.DB
	query      *query.Query
	serviceFun xgorm.IServiceFunctions[model.SysUser]
}

func (s SysUserService) AddUser(c *gin.Context, param *AddOrUpdateSysUserParam) (err error) {
	menuQuery := s.query.SysUser
	insertC := &model.SysUser{
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
	}
	s.serviceFun.BindCreateInfo(c, insertC)
	err = menuQuery.WithContext(c).Create(insertC)
	return
}

func (s SysUserService) UpdateUser(c *gin.Context, id int32, param *AddOrUpdateSysUserParam) (err error) {
	menuQuery := s.query.SysUser
	updateC := &model.SysUser{
		Nickname:       param.NickName,
		Username:       param.UserName,
		Roles:          strings.Join(param.UserRoles, ","),
		Phone:          param.UserPhone,
		Avatar:         "",
		Email:          param.UserEmail,
		Gender:         param.UserGender,
		LastOnlineTime: time.Time{},
		LastCpwdTime:   time.Time{},
		UserStatus:     param.Status,
		Version:        0,
	}
	s.serviceFun.BindUpdateInfo(c, updateC)
	updates, err := menuQuery.WithContext(c).
		Where(menuQuery.ID.Eq(id)).
		Where(menuQuery.DeleteTag.Eq(0)).
		Updates(updateC)
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
	operatorInfo := s.serviceFun.GetOperatorInfo(c)
	info, err := menuQuery.WithContext(c).Where(menuQuery.ID.In(ids...)).Updates(&model.SysUser{
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
	if info.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_AFFECT_NONE_ERROR)
	}
	return nil
}

func (s SysUserService) SysUserList(c *gin.Context, param *SysUserListParam) (sysUserListResp SysUserListResp, err error) {
	sysUserQuery := query.Use(s.db).SysUser
	sysUserListResp.PageResult.PageParam = param.PageParam
	userListInSql, totalCount, err := sysUserQuery.WithContext(c).
		Where(sysUserQuery.DeleteTag.Eq(0)).
		FindByPage((param.Current-1)*param.Size, param.Size)
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
