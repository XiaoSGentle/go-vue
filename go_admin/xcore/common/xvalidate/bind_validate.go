package xvalidate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"xcore/common/xerror"
)

// ValidateStruct  用于绑定地址栏参数
func ValidateStruct(param interface{}) (err error) {
	if thisTranslator == nil || thisValidate == nil {
		err = xerror.NewErrCode(xerror.SERVER_COMMON_ERROR)
		return
	}
	//获取验证器
	if _err := thisValidate.Struct(param); _err != nil {
		var errs validator.ValidationErrors
		errors.As(_err, &errs)
		// errs.Translate(thisTranslator) 返回的是一个map[string]string  获取后返回第一个就好了
		err = xerror.NewErrCodeMsg(xerror.REUQEST_PARAM_ERROR, getFirstKeyValue(errs.Translate(thisTranslator)))
		return
	}
	err = nil
	return
}
func getFirstKeyValue(m map[string]string) string {
	for _, value := range m {
		return value
	}
	return ""
}
