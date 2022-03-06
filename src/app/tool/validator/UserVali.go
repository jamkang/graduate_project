package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//电话验证
func checkTele(f1 validator.FieldLevel) bool {
	data, _ := f1.Field().Interface().(string)
	reg := regexp.MustCompile(`^1\d{10}$`)
	if la := reg.MatchString(data); !la {
		return false
	}
	return true
}
