package xcron

import (
	"fmt"
	"github.com/robfig/cron"
	"reflect"
	"strings"
)

type IXCron interface {
	ScannerFunctions() ([]string, []ScannerFunctionResult)
	AddScannerModule(interface{})
	SetCronRule(func(*cron.Cron))
	CallMethodsOnValue(key string, args ...interface{})
	StartCorn()
	ReStartCorn()
}
type XCron struct {
	scannerModule []interface{}
	mapModule     map[string]interface{}
	cron          *cron.Cron
	resetFunc     func()
	lastChan      chan struct{}
}

func (c *XCron) SetCronRule(f func(*cron.Cron)) {
	entries := c.cron.Entries()
	entries = append(entries[:0], entries[len(entries):]...)
	c.resetFunc = func() {
		f(c.cron)
	}
}

func (c *XCron) StartCorn() {
	c.resetFunc()
	go c.cron.Run()

}
func (c *XCron) ReStartCorn() {
	c.cron.Stop()
	c.cron = cron.New()
	c.resetFunc()
	go c.cron.Run()
}

type ScannerFunctionResult struct {
	Key        string
	ParamTypes []string
}

func (c *XCron) AddScannerModule(i interface{}) {
	c.scannerModule = append(c.scannerModule, i)
}

func NewXCorn(scannerModule ...interface{}) IXCron {
	return &XCron{
		cron:          cron.New(),
		scannerModule: scannerModule,
		mapModule:     make(map[string]interface{}),
	}
}

func (c *XCron) ScannerFunctions() (key []string, results []ScannerFunctionResult) {
	// 从接口值获取反射值
	reflectValue := reflect.ValueOf(c.scannerModule)
	// 遍历切片元素
	for i := 0; i < reflectValue.Len(); i++ {
		elem := reflectValue.Index(i)
		// pull进值
		var s = reflect.ValueOf(elem.Interface()).Type().Name()
		var e = elem.Interface()
		c.mapModule[s] = e
		// s
		results = append(results, c.scannerFunction(elem.Interface())...)
	}

	return c.extractKeys(results), results
}
func (c *XCron) scannerFunction(value interface{}) []ScannerFunctionResult {

	var results []ScannerFunctionResult
	// 从接口值获取反射值
	reflectValue := reflect.ValueOf(value)

	// 如果传入的是指针,获取底层值
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	// 获取值的类型
	typ := reflectValue.Type()

	// 遍历类型的方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)

		// 获取方法的值
		methodValue := reflectValue.Method(i)

		// 检查方法是否有效
		if !methodValue.IsValid() {
			continue
		}
		var functionNeedTypes []string
		// 打印方法需要的参数类型

		for j := 1; j < method.Type.NumIn(); j++ {
			paramType := method.Type.In(j)
			functionNeedTypes = append(functionNeedTypes, paramType.String())
		}

		results = append(results, ScannerFunctionResult{
			Key:        fmt.Sprintf("%s::%s", typ.Name(), method.Name),
			ParamTypes: functionNeedTypes,
		})
	}
	return results
}
func (c *XCron) CallMethodsOnValue(key string, args ...interface{}) {
	funcKey := strings.Split(key, "::")
	// 从接口值获取反射值
	reflectValue := reflect.ValueOf(c.mapModule[funcKey[0]])
	// 如果传入的是指针,获取底层值
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	// 获取值的类型
	typ := reflectValue.Type()

	// 遍历类型的方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		if method.Name == funcKey[1] {
			// 获取方法的值
			methodValue := reflectValue.Method(i)

			// 检查方法是否有效
			if !methodValue.IsValid() {
				continue
			}

			// 构造方法参数值
			methodArgs := make([]reflect.Value, len(args))
			for j, arg := range args {
				methodArgs[j] = reflect.ValueOf(arg)
			}
			// 调用方法
			results := methodValue.Call(methodArgs)

			var cronFunResult string
			// 打印结果
			for _, result := range results {
				cronFunResult += fmt.Sprintf("%s", result.Interface())
			}
		}
	}

}
func (c *XCron) extractKeys(results []ScannerFunctionResult) []string {
	keys := make([]string, len(results))
	for i, result := range results {
		keys[i] = result.Key
	}
	return keys
}
