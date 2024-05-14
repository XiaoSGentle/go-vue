package xstring

import (
	"regexp"
	"strings"
)

func SnakeToLowerCamelCase(columnName string) string {
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
}
func LowerCamelCaseToSnake(columnName string) string {

	// 使用正则表达式将大写字母前面插入下划线，并将所有字母转换为小写
	re := regexp.MustCompile(`(.)([A-Z][a-z]+)`)
	converted := re.ReplaceAllString(columnName, "${1}_${2}")

	// 将所有字母转换为小写
	converted = strings.ToLower(converted)

	return converted

}
func SnakeToUpperCamelCase(columnName string) string {
	lowUpStr := SnakeToLowerCamelCase(columnName)
	upUpStr := strings.Title(lowUpStr)
	return upUpStr
}
func GoTypeConversion(columnType string) string {
	m := map[string]string{
		"varchar": "string",
		"int":     "int",
	}
	goLangType := m[columnType]
	if goLangType == "" {
		goLangType = "string"
	}
	return goLangType
}
func TsTypeConversion(columnType string) string {
	m := map[string]string{
		"varchar": "string",
		"int":     "number",
	}
	goLangType := m[columnType]
	if goLangType == "" {
		goLangType = "string"
	}
	return goLangType
}
