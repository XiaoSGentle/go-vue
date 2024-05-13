package sys_cron

import (
	"xcore/common/xtype/xbase"
)

type SysCronListResp struct {
	xbase.PageResult
	Records []CronResp `json:"records"`
}

type CronResp struct {
	Key         string   `json:"key" zh_comment:"标识" en_comment:"key" validate:"required"`                 // 唯一键值
	Description string   `json:"description" zh_comment:"描述" en_comment:"description" validate:"required"` // 描述
	Schedule    string   `json:"schedule" zh_comment:"cron表达式" en_comment:"cron code" validate:"required"` // cron表达式
	Status      string   `json:"status" zh_comment:"状态" en_comment:"status" validate:"required"`           // 状态
	Arguments   []string `json:"arguments" zh_comment:"参数" en_comment:"arguments" `                        // 参数
}

type UpdateCronReq struct {
	Key         string   `json:"key" zh_comment:"标识" en_comment:"key" validate:"required"`                 // 唯一键值
	Description string   `json:"description" zh_comment:"描述" en_comment:"description" validate:"required"` // 描述
	Schedule    string   `json:"schedule" zh_comment:"cron表达式" en_comment:"cron code" validate:"required"` // cron表达式
	Status      string   `json:"status" zh_comment:"状态" en_comment:"status" validate:"required"`           // 状态
	Arguments   []string `json:"arguments" zh_comment:"参数" en_comment:"arguments" `                        // 参数
}
