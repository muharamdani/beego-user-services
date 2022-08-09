package utils

import (
	"github.com/beego/beego/v2/core/validation"
	"strings"
)

func GetValidationResult(model interface{}) map[string]string {
	validationResult := make(map[string]string)
	valid := validation.Validation{}
	b, _ := valid.Valid(model)
	if !b {
		for _, err := range valid.Errors {
			key := strings.Split(err.Key, ".")
			validationResult[strings.ToLower(key[0])] = err.Message
		}
	}
	return validationResult
}
