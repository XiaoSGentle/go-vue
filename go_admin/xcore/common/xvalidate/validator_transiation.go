package xvalidate

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var thisValidate *validator.Validate
var thisTranslator ut.Translator

func InitTransValidator(lang string) {
	//设置支持语言
	chinese := zh.New()
	english := en.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese, english)
	//设置验证器
	thisValidate = validator.New()
	//根据参数取翻译器实例
	thisTranslator, _ = uni.GetTranslator(lang)
	//翻译器注册到validator
	switch lang {
	case "chinese":
		err := zhTranslations.RegisterDefaultTranslations(thisValidate, thisTranslator)
		if err != nil {
			return
		}
		//使用fld.Tag.Get("comment")注册一个获取tag的自定义方法
		thisValidate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("zh_comment")
		})
	case "english":
		err := enTranslations.RegisterDefaultTranslations(thisValidate, thisTranslator)
		if err != nil {
			return
		}
		thisValidate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en_comment")
		})
	}
}
