package sys_gen

type GenTableParam struct {
	Names []string `json:"names" zh_comment:"表名" en_comment:"table names" validate:"required,min=1"`
}
type GenTableColumnsParam struct {
	Name string `uri:"name" zh_comment:"表名" en_comment:"table names" validate:"required,min=1"`
}

type GenTablePreview struct {
	Lang        string `json:"lang,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	FileContent string `json:"fileContent,omitempty"`
}
