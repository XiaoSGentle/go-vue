package sys_dict

import (
	"github.com/gin-gonic/gin"
	baseType "xadmin/soybean/dao/model/base"
	"xcore/common/xcache"
	"xcore/common/xerror"
	"xcore/common/xmiddlewares"
	"xcore/common/xresponse"
	"xcore/common/xvalidate"
	"xcore/core/xcore"
)

var SysDictGroup = xcore.Group("/system/dict", newSysMenuHandler, regSysDict, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize, xmiddlewares.Verify)
var SysDictDataGroup = xcore.Group("/system/dict/data", newSysMenuHandler, regSysDictData, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize, xmiddlewares.Verify)
var NoAuthDictGroup = xcore.Group("/dict", newSysMenuHandler, regNoAuthSysDict, xmiddlewares.LogMiddleHandler, xmiddlewares.Authorize)
var DictCache = xcache.NewCacheStore[[]DictInfo](0, "DICT_CACHE_PREFIX_KEY::")

func regSysDict(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysDictHandler) {
		rg.GET("/list", handle.GetDictList)
		rg.POST("", handle.AddDict)
		rg.PUT("/:id", handle.UpdateDict)
		rg.DELETE("", handle.DeleteDictByIds)
	})
}
func regNoAuthSysDict(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysDictHandler) {
		rg.GET("/:typeCode", handle.DictInfo)
	})
}

func regSysDictData(rg *gin.RouterGroup, group *xcore.GroupBase) error {
	return group.Reg(func(handle *sysDictHandler) {
		rg.GET("/list", handle.GetDictDataList)
		rg.POST("", handle.AddDictData)
		rg.PUT("/:id", handle.UpdateDictData)
		rg.DELETE("", handle.DeleteDictDataByIds)
	})
}

type sysDictHandler struct {
	dictService     ISysDictService
	dictDataService ISysDictDataService
}

func newSysMenuHandler(dict ISysDictService, dictData ISysDictDataService) *sysDictHandler {
	return &sysDictHandler{dictService: dict, dictDataService: dictData}
}

func (h sysDictHandler) GetDictList(c *gin.Context) {
	var param SysDictListParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	list, err := h.dictService.GetDictList(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)

}

func (h sysDictHandler) AddDict(c *gin.Context) {
	var sysMenuAddParam AddOrUpDateSysDictParam
	if err := c.ShouldBind(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	err := h.dictService.AddDict(c, &sysMenuAddParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}

func (h sysDictHandler) UpdateDict(c *gin.Context) {
	var param AddOrUpDateSysDictParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	var pathId baseType.DetailsId
	err := c.BindUri(&pathId)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err = h.dictService.UpdateDict(c, pathId.Id, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.UpdateSuccessCtx(c)
}

func (h sysDictHandler) DeleteDictByIds(c *gin.Context) {
	var ids baseType.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := h.dictService.DeleteDict(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}

func (h sysDictHandler) GetDictDataList(c *gin.Context) {
	var param SysDictDataListParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	list, err := h.dictDataService.GetDictDataList(c, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.SuccessCtx(c, list)

}

func (h sysDictHandler) AddDictData(c *gin.Context) {
	var sysMenuAddParam AddOrUpDateSysDictDataParam
	if err := c.ShouldBind(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&sysMenuAddParam); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}

	err := h.dictDataService.AddDictData(c, &sysMenuAddParam)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.CreateSuccessCtx(c)
}

func (h sysDictHandler) UpdateDictData(c *gin.Context) {
	var param AddOrUpDateSysDictDataParam
	if err := c.ShouldBind(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := xvalidate.ValidateStruct(&param); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	var pathId baseType.DetailsId
	err := c.BindUri(&pathId)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	err = h.dictDataService.UpdateDictData(c, pathId.Id, &param)
	if err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.UpdateSuccessCtx(c)
}

func (h sysDictHandler) DeleteDictDataByIds(c *gin.Context) {
	var ids baseType.DelIds
	if err := c.ShouldBind(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if err := xvalidate.ValidateStruct(&ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	if err := h.dictDataService.DeleteDictData(c, ids.Ids); err != nil {
		xresponse.ErrorCtx(c, err)
		return
	}
	xresponse.DeleteSuccessCtx(c)
}

func (h sysDictHandler) DictInfo(c *gin.Context) {
	var typeCode struct {
		TypeCode string `json:"typeCode" zh_comment:"字典KEY" en_comment:"dict key" validate:"required"`
	}
	if err := c.ShouldBind(&typeCode); err != nil {
		xresponse.ErrorCtx(c, err)
	}
	if !DictCache.Exist(typeCode.TypeCode) {
		for key, value := range h.DictInSql(c) {
			DictCache.Set(key, value)
		}
		if !DictCache.Exist(typeCode.TypeCode) {
			xresponse.ErrorCtx(c, xerror.NewErrCode(xerror.DICT_NOT_EXIST_ERROR))
			return
		} else {
			xresponse.SuccessCtx(c, DictCache.Get(typeCode.TypeCode))
		}
	} else {
		xresponse.SuccessCtx(c, DictCache.Get(typeCode.TypeCode))
	}
	return
}

func (h sysDictHandler) DictInSql(c *gin.Context) (result map[string][]DictInfo) {
	list, _ := h.dictService.GetDictList(c, &SysDictListParam{
		PageParam: baseType.PageParam{
			Current: 1,
			Size:    1000,
		},
	})
	for _, record := range list.Records {
		dataList, _ := h.dictDataService.GetDictDataList(c, &SysDictDataListParam{
			PageParam: baseType.PageParam{
				Current: 1,
				Size:    1000,
			},
			Code: record.Code,
		})
		var dictInfo []DictInfo
		for _, dictDataList := range dataList.Records {
			dictInfo = append(dictInfo, DictInfo{
				Label: dictDataList.Label,
				Value: dictDataList.Value,
				Sort:  dictDataList.Sort,
			})
		}
		result[record.Code] = dictInfo
	}
	return result
}
