package xcore

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"xcore/core/xdig"
)

type GroupBase struct {
	basePath    string                                            // 基础路径
	initHandle  interface{}                                       //？
	regHandle   func(rg *gin.RouterGroup, group *GroupBase) error //？ 两个参数，ginRouter 组 自定义路由组
	middlewares []gin.HandlerFunc                                 // gin使用的中间件
}

// Group 创建一个路由组
func Group(relativePath string, initHandle interface{}, regHandle func(rg *gin.RouterGroup, group *GroupBase) error,
	middlewares ...gin.HandlerFunc) *GroupBase {
	return &GroupBase{
		basePath:    relativePath,
		initHandle:  initHandle,
		regHandle:   regHandle,
		middlewares: middlewares,
	}
}

// RegisterGroup 将路由组注册到gin
func RegisterGroup(rg *gin.RouterGroup, group *GroupBase) {
	r := rg.Group(group.basePath)
	if len(group.middlewares) > 0 {
		r.Use(group.middlewares...)
	}

	if err := xdig.ProvideForDI(group.initHandle); err != nil {
		log.Fatalln(err)
	}
	if err := group.regHandle(r, group); err != nil {
		log.Fatalln(err)
	}
}

// Reg registers handle by DI
func (group GroupBase) Reg(function interface{}, opts ...dig.InvokeOption) error {
	return xdig.DI(function, opts...)
}
