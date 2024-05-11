package xgorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"time"
	"xcore/common/xerror"
	"xcore/common/xtoken"
	"xcore/common/xtype/xbase"
)

type IServiceFunctions[T interface{}] interface {
	FindOneById(c *gin.Context, id int32) (error, T)
	FindByPage(c *gin.Context, param xbase.PageParam, conditions ...func(*gorm.DB) *gorm.DB) (error, xbase.PageResultList[T])
	UpDateById(c *gin.Context, id int32, entity T) (error, int32)
	DeleteByIds(c *gin.Context, ids []int32) (error, int32)
	SoftDeleteByIds(c *gin.Context, ids []int32) (error, int32)
	Exist(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) bool
	InsertOne(c *gin.Context, entity T) (error, int32)
	InsertBatch(c *gin.Context, entities []T) (error, int32)
	Count(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) int32
}

func InjectService[T interface{}](db *gorm.DB) IServiceFunctions[T] {
	return &ServiceFunction[T]{
		db: db,
	}
}

type ServiceFunction[T interface{}] struct {
	db *gorm.DB
}

func (i ServiceFunction[T]) FindOneById(c *gin.Context, id int32) (error, T) {
	result, db := i.getInstanceOfTAndDb(c)
	if !i.Exist(c, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}) {
		return xerror.NewErrCode(xerror.CURD_DATA_NOT_EXIST_ERROR), result
	}
	db.Where("id = ?", id).Find(&result)
	return nil, result
}
func (i ServiceFunction[T]) FindByPage(c *gin.Context, param xbase.PageParam, conditions ...func(*gorm.DB) *gorm.DB) (error, xbase.PageResultList[T]) {
	var result xbase.PageResultList[T]
	result.PageParam = param
	resultList, db := i.getInstanceListOfTAndDbAddition(c, conditions...)
	tx := db.Offset((param.Current - 1) * param.Size).Limit(param.Size).Find(&resultList)
	if tx.Error != nil {
		return tx.Error, result
	}
	result.Records = resultList
	if size := len(resultList); 0 < param.Size && 0 < size && size < param.Size {
		result.Total = int64(size + (param.Current-1)*param.Size)
		return nil, result
	}
	tx = db.Offset(-1).Limit(-1).Count(&result.Total)
	if tx.Error != nil {
		return tx.Error, result
	}
	return nil, result
}

func (i ServiceFunction[T]) UpDateById(c *gin.Context, id int32, entity T) (error, int32) {
	db := i.getInstanceOfDB(c)
	i.BindUpdateInfo(c, &entity)
	tx := db.Where("id = ?", id).Updates(&entity)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_UPDATE_AFFECT_NONE_ERROR), 0
	}
	return nil, int32(tx.RowsAffected)
}

func (i ServiceFunction[T]) DeleteByIds(c *gin.Context, ids []int32) (error, int32) {
	instance, db := i.getInstanceListOfTAndDb(c)

	tx := db.Unscoped().Delete(&instance, ids)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_UPDATE_AFFECT_NONE_ERROR), 0
	}
	return nil, int32(tx.RowsAffected)
}

func (i ServiceFunction[T]) SoftDeleteByIds(c *gin.Context, ids []int32) (error, int32) {
	db := i.getInstanceOfDB(c)
	tx := db.Where("id IN ?", ids).Updates(map[string]interface{}{
		"soft_delete_tag": "1",
		"update_time":     time.Now(),
		"update_uid":      i.GetOperatorInfo(c).Uid,
		"update_by":       i.GetOperatorInfo(c).NickName,
	})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_UPDATE_AFFECT_NONE_ERROR), 0
	}
	return nil, int32(tx.RowsAffected)
}

func (i ServiceFunction[T]) Exist(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) bool {
	return i.Count(c, conditions...) > 0
}

func (i ServiceFunction[T]) InsertOne(c *gin.Context, entity T) (error, int32) {
	db := i.getInstanceOfDB(c)
	i.BindCreateInfo(c, &entity)
	tx := db.Create(&entity)
	if tx.Error != nil {
		return tx.Error, 0
	} else if tx.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_UPDATE_AFFECT_NONE_ERROR), 0
	}
	return nil, int32(tx.RowsAffected)
}

func (i ServiceFunction[T]) InsertBatch(c *gin.Context, entities []T) (error, int32) {
	_, db := i.getInstanceListOfTAndDb(c)
	i.BindCreateInfo(c, &entities)
	tx := db.Create(&entities)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return xerror.NewErrCode(xerror.CURD_UPDATE_AFFECT_NONE_ERROR), 0
	}
	return nil, int32(tx.RowsAffected)
}

func (i ServiceFunction[T]) Count(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) int32 {
	count := int64(0)
	_, db := i.getInstanceOfTAndDbAddition(c, conditions...)
	db.Count(&count)
	return int32(count)
}

func (i ServiceFunction[T]) GetInstanceOfT() (result T) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return *thisResult
}
func (i ServiceFunction[T]) getInstanceOfDB(c *gin.Context) (db *gorm.DB) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return i.db.Model(thisResult).WithContext(c)
}
func (i ServiceFunction[T]) getInstanceOfListOfT() (result []T) {
	sliceType := reflect.SliceOf(reflect.TypeOf((*T)(nil)).Elem())
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	result = sliceValue.Interface().([]T)
	return
}
func (i ServiceFunction[T]) getInstanceOfTAndDb(c *gin.Context) (result T, db *gorm.DB) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return *thisResult, i.db.Model(thisResult).WithContext(c)
}
func (i ServiceFunction[T]) getInstanceListOfTAndDb(c *gin.Context) (result []T, db *gorm.DB) {
	sliceType := reflect.SliceOf(reflect.TypeOf((*T)(nil)).Elem())
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	return sliceValue.Interface().([]T), i.db.Model(thisResult).WithContext(c)
}
func (i ServiceFunction[T]) getInstanceOfTAndDbAddition(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) (result T, db *gorm.DB) {
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	contextDb := i.db.Model(thisResult).WithContext(c)
	for _, condition := range conditions {
		contextDb = condition(contextDb)
	}
	return *thisResult, contextDb
}
func (i ServiceFunction[T]) getInstanceListOfTAndDbAddition(c *gin.Context, conditions ...func(*gorm.DB) *gorm.DB) (result []T, db *gorm.DB) {
	sliceType := reflect.SliceOf(reflect.TypeOf((*T)(nil)).Elem())
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	ptrValue := reflect.New(reflect.TypeOf((*T)(nil)).Elem())
	thisResult := ptrValue.Elem().Addr().Interface().(*T)
	contextDb := i.db.Model(thisResult).WithContext(c)
	for _, condition := range conditions {
		contextDb = condition(contextDb)
	}
	return sliceValue.Interface().([]T), contextDb
}
func (i ServiceFunction[T]) BindCreateInfo(c *gin.Context, target interface{}) {
	BindCreateInfo(target, c)
}
func (i ServiceFunction[T]) BindUpdateInfo(c *gin.Context, target interface{}) {
	BindUpdateInfo(target, c)
}
func (i ServiceFunction[T]) GetOperatorInfo(c *gin.Context) *xtoken.ClaimsPayload {
	return xtoken.GetBindCustomPayload(c)
}
