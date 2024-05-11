package xgorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"xcore/common/xresponse"
	"xcore/common/xtype/xbase"
	"xcore/common/xvalidate"
)

type IRouterFunctions[T interface{}] interface {
	FindOneById(c *gin.Context)
	FindByPage(c *gin.Context, queryHelper func(IQueryHelper) *gorm.DB)
	UpDateById(c *gin.Context)
	DeleteByIds(c *gin.Context)
	SoftDeleteByIds(c *gin.Context)
	Create(c *gin.Context)
}

func InjectRouter[T interface{}](db *gorm.DB) IRouterFunctions[T] {
	return &RouterFunction[T]{
		serviceFunctions: InjectService[T](db),
	}
}

type RouterFunction[T interface{}] struct {
	serviceFunctions IServiceFunctions[T]
}

func (r RouterFunction[T]) FindOneById(c *gin.Context) {
	var id xbase.DetailsId
	if err := c.ShouldBindUri(&id); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&id); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err, t := r.serviceFunctions.FindOneById(c, id.Id)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, t)
}

func (r RouterFunction[T]) FindByPage(c *gin.Context, queryHelper func(IQueryHelper) *gorm.DB) {
	var pageParam xbase.PageParam
	if err := c.ShouldBind(&pageParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&pageParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err, result := r.serviceFunctions.FindByPage(c, pageParam, func(db *gorm.DB) *gorm.DB {
		return queryHelper(NewQueryHelper(db))
	})
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, result)
}

func (r RouterFunction[T]) UpDateById(c *gin.Context) {
	param := r.getInstanceOfT()
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	var pathId xbase.DetailsId
	err := c.BindUri(&pathId)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err = xvalidate.ValidateStruct(&pathId); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err, _ = r.serviceFunctions.UpDateById(c, pathId.Id, param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
	}
	xresponse.UpdateSuccessCtx(c)

}

func (r RouterFunction[T]) DeleteByIds(c *gin.Context) {
	var ids xbase.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err, _ := r.serviceFunctions.DeleteByIds(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}

func (r RouterFunction[T]) SoftDeleteByIds(c *gin.Context) {
	var ids xbase.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err, _ := r.serviceFunctions.SoftDeleteByIds(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}

func (r RouterFunction[T]) Create(c *gin.Context) {
	param := r.getInstanceOfT()
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err, _ := r.serviceFunctions.InsertOne(c, param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}
func (r RouterFunction[T]) getInstanceOfT() (result T) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return *thisResult
}
