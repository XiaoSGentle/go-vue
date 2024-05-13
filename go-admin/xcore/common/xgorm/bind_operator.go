package xgorm

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
	"xcore/common/xtoken"
	"xcore/core/xvariable"
)

// BindCreateInfo  这个工具是填充 插入数据时候的添加用户的信息
func BindCreateInfo(target interface{}, c *gin.Context) {
	if t := reflect.TypeOf(target); t.Kind() != reflect.Ptr {
		xvariable.Logger.ErrorLog.ErrorContext(c, "传入需要绑定的目标必须是一个指针类型")
	}
	operateTime := time.Now()
	operateInfo := xtoken.GetBindCustomPayload(c)
	needSetData := map[string]interface{}{
		"CreateTime": operateTime,
		"CreateBy":   operateInfo.NickName,
		"CreateUID":  operateInfo.Uid,
	}
	// 数组类型
	if t := reflect.TypeOf(target); t.Kind() == reflect.Array {
		// 如果 target 是数组类型
		arr := reflect.ValueOf(target)
		for i := 0; i < arr.Len(); i++ {
			_target := arr.Index(i)
			// 对数组中的每个元素执行操作
			bindInfo(needSetData, _target)
		}
		return
	}
	bindInfo(needSetData, target)
}

// BindUpdateInfo  这个工具是填充 更新数据时候的添加用户的信息
func BindUpdateInfo(target interface{}, c *gin.Context) {
	if t := reflect.TypeOf(target); t.Kind() != reflect.Ptr {
		xvariable.Logger.ErrorLog.ErrorContext(c, "传入需要绑定的目标必须是一个指针类型")
	}
	operateTime := time.Now()
	operateInfo := xtoken.GetBindCustomPayload(c)
	needSetData := map[string]interface{}{
		"UpdateTime": operateTime,
		"UpdateUID":  operateInfo.Uid,
		"UpdateBy":   operateInfo.NickName,
	}
	bindInfo(needSetData, target)
}

func bindInfo(bindInfo map[string]interface{}, target interface{}) {
	_target := reflect.ValueOf(target).Elem()
	for setKey, setValue := range bindInfo {
		if value := _target.FieldByName(setKey); value.CanSet() {
			value.Set(reflect.ValueOf(setValue))
		}
	}
}
