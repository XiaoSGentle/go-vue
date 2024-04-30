package xdig

import (
	"go.uber.org/dig"
	"reflect"
)

// 创建一个dig依赖注入容器
var container = dig.New()
var instanceList []interface{}

// ProvideForDI 提供依赖
func ProvideForDI(constructor interface{}, opts ...dig.ProvideOption) error {
	// 防止重复注册报错
	for _, instance := range instanceList {
		if reflect.ValueOf(instance).String() == reflect.ValueOf(constructor).String() {
			return nil
		}
	}
	instanceList = append(instanceList, constructor)
	return container.Provide(constructor, opts...)

}

// DI 生成目标

func DI(function interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}
