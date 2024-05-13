package xvalidate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"reflect"
	"xcore/common/xerror"
)

var BIND_GIN_CONTEXT_KEY = "BIND_GIN_CONTEXT_KEY"

type IValidator interface {
	ValidateStruct(param interface{}) (err error)
}

type Validator struct {
	selectTranslator ut.Translator
	selectValidator  *validator.Validate

	chineseTranslator ut.Translator
	englishTranslator ut.Translator
	chineseValidator  *validator.Validate
	englishValidator  *validator.Validate
}

func NewValidator() IValidator {

	chineseDict := zh.New()
	englishDict := en.New()
	//设置国际化翻译器
	uni := ut.New(chineseDict, englishDict, englishDict)
	//设置验证器
	chineseValidator := validator.New()
	englishValidator := validator.New()
	//根据参数取翻译器实例
	chineseTranslator, _ := uni.GetTranslator("zh")
	englishTranslator, _ := uni.GetTranslator("en")
	//翻译器注册到validator

	err := zhTranslations.RegisterDefaultTranslations(chineseValidator, chineseTranslator)

	if err != nil {

		panic(err)
		return nil
	}
	chineseValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("zh_comment")
	})
	err = enTranslations.RegisterDefaultTranslations(englishValidator, englishTranslator)
	if err != nil {

		panic(err)
		return nil
	}
	englishValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("en_comment")
	})
	return &Validator{
		selectTranslator:  chineseTranslator,
		selectValidator:   chineseValidator,
		chineseTranslator: chineseTranslator,
		englishTranslator: englishTranslator,
		chineseValidator:  chineseValidator,
		englishValidator:  englishValidator,
	}
}

// ValidateStruct  用于绑定地址栏参数
func (v *Validator) ValidateStruct(param interface{}) (err error) {
	if v.selectValidator == nil {
		err = xerror.NewErrCode(xerror.SERVER_COMMON_ERROR)
		return
	}
	//获取验证器
	if _err := v.selectValidator.Struct(param); _err != nil {
		var errs validator.ValidationErrors
		errors.As(_err, &errs)
		// errs.Translate(thisTranslator) 返回的是一个map[string]string  获取后返回第一个就好了
		err = xerror.NewErrCodeMsg(xerror.REUQEST_PARAM_ERROR, getFirstKeyValue(errs.Translate(v.selectTranslator)))
		return
	}
	err = nil
	return
}

func (v *Validator) ChangeEnglishTranslator() {
	v.selectTranslator = v.englishTranslator
	v.selectValidator = v.englishValidator

}
func (v *Validator) ChangeChineseTranslator() {
	v.selectValidator = v.chineseValidator
	v.selectTranslator = v.chineseTranslator
}
func getFirstKeyValue(m map[string]string) string {
	for _, value := range m {
		return value
	}
	return ""
}

func GetBindValidator(c *gin.Context) IValidator {
	value, exists := c.Get(BIND_GIN_CONTEXT_KEY)
	if exists && value != nil {
		return value.(IValidator)
	}
	return NewValidator()
}
