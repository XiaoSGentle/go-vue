package sys_gen

type GenTableParam struct {
	Names []string `json:"names" zh_comment:"表名" en_comment:"table names" validate:"required,min=1"`
}
type GenTableColumnsParam struct {
	Name string `uri:"name" zh_comment:"表名" en_comment:"table names" validate:"required,min=1"`
}

type GenTablePreview struct {
	GoRouter     string `json:"goRouter,omitempty"`
	GoService    string `json:"goService,omitempty"`
	GoType       string `json:"goType,omitempty"`
	VueIndex     string `json:"vueIndex,omitempty"`
	VueDrawer    string `json:"vueDrawer,omitempty"`
	VueApi       string `json:"vueApi,omitempty"`
	VueNameSpace string `json:"vueNameSpace,omitempty"`
}
