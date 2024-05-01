package base_type

type PageParam struct {
	Current int `json:"current" form:"current" zh_comment:"当前页数" en_comment:"current" validate:"required,gte=1"` // 必填，页面值>=1
	Size    int `json:"size" form:"size" zh_comment:"每页条数" en_comment:"size" validate:"required,gte=1"`          // 必填，每页条数值>=1
}

type BaseRecord struct {
	/** record Id */
	ID int32 `json:"id"`
	/** record creator */
	CreateBy string `json:"createBy"`
	/** record create time */
	CreateTime string `json:"createTime"`
	/** record updater */
	UpdateBy string `json:"updateBy"`
	/** record update time */
	UpdateTime string `json:"updateTime"`
	/** record status */
	Status string `json:"status"`
}

type PageResult struct {
	PageParam
	Total int64 `json:"total"`
}
type DelIds struct {
	Ids string `uri:"ids"`
}

type DetailsId struct {
	Id string `uri:"id"`
}
