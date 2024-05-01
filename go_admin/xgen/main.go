package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strings"
)

var (
	User     = "xiaos_core_new"
	Pass     = "5SJPJXbMFbyJEKp5"
	Host     = "xiaos.cool"
	Port     = "3306"
	DataBase = "xiaos_core_new"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath:           "../xcore/dao/query",
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", User, Pass, Host, Port, DataBase)

	mysqlDialectic := mysql.Open(connectUrl)

	gormDb, err := gorm.Open(mysqlDialectic, &gorm.Config{})
	if err != nil {
		return
	}
	g.UseDB(gormDb)
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int32" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return tag.Set("autoCreateTime", "milli")
	})
	autoUpdateTimeFiled := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return tag.Set("autoUpdateTime", "milli")
	})
	//softDeleteField := gen.FieldType("soft_delete_tag", "gorm.DeletedAt")
	// 下划线转驼峰
	g.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		// 去掉下划线和横杠，并将后面的字母改为大写
		var modifiedName string
		for i := 0; i < len(columnName); i++ {
			if columnName[i] == '_' {
				if i+1 < len(columnName) {
					modifiedName += strings.ToUpper(string(columnName[i+1]))
					i++
				}
			} else {
				modifiedName += string(columnName[i])
			}
		}

		return modifiedName
	})
	// 创建模型的结构体,生成文件在 module 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	//Address := g.GenerateModel("address")
	//ignore := gen.FieldIgnore("uuid", "create_name", "create_time", "create_uuid", "update_time", "update_uuid",
	//	"update_name", "soft_delete_tag", "version")
	// 创建全部模型文件, 并覆盖前面创建的同名模型

	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeFiled}
	allModel := g.GenerateAllTable(fieldOpts...)
	// // 创建有关联关系的模型文件
	//physiological := g.GenerateModel("physiological",
	//	append(
	//		fieldOpts,
	//		// user 一对多 address 关联, 外键`uid`在 address 表中
	//		gen.WithMethod(Physiological{}),
	//	)...,
	//)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	//g.ApplyBasic(physiological)
	g.ApplyBasic(allModel...)
	g.Execute()

}
