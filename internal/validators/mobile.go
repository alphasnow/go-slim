package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func MobileNumberValidation(fl validator.FieldLevel) bool {
	if mobile, ok := fl.Field().Interface().(string); ok {
		if len(mobile) != 11 {
			return false
		}
	}
	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile_number", MobileNumberValidation)
	}
}
