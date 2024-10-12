package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldsErrors map[string]string
}

// return true if fieldError
func (v *Validator) HasError() bool {
	return len(v.FieldsErrors) != 0
}

// add eror
func (v *Validator) AddError(key string, message string) {
	if v.FieldsErrors == nil {
		v.FieldsErrors = make(map[string]string)
	}

	if _, exists := v.FieldsErrors[key]; !exists {
		v.FieldsErrors[key] = message
	}
}

// check fieldsError
func (v *Validator) CheckField(err bool, key string, message string) {
	if err {
		v.AddError(key, message)
	}
}

func IsStringEmpty(message string) bool {
	return strings.TrimSpace(message) == ""
}

func MaxChar(value string, n int) bool {
	return utf8.RuneCountInString(value) > n
}

// return true if the value isn't permited
func NotPermitedInt(value int, permitedValues ...int) bool {
	for i := range value {
		if value == permitedValues[i] {
			return false
		}
	}
	return true
}
