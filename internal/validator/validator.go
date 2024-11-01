package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	NonFieldErrors []string
	FieldsErrors   map[string]string
}

// return true if fieldError
func (v *Validator) HasError() bool {
	return len(v.FieldsErrors) != 0 && len(v.NonFieldErrors) != 0
}

// adds messge to the nonfieldErros
func (v *Validator) AddNonFieldErros(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
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

// add erros if when err it true
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

var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// MinChars() returns true if a value contains at least n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Matches() returns true if a value matches a provided compiled regular
// expression pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
