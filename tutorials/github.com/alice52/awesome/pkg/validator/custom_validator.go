package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 可以使用 fl.Field() 获取字段值进行验证 + 返回 true 表示验证通过，返回 false 表示验证失败
func numberTag(fl validator.FieldLevel) bool {
	matchPattern := "^[0-9]+$"
	fieldValue := fl.Field().String()
	matched, _ := regexp.MatchString(matchPattern, fieldValue)
	return matched
}

func register(c *validator.Validate, tagName string, vlf validator.Func) *validator.Validate {

	// 自定义标签 "customTag"
	err := c.RegisterValidation(tagName, vlf)
	if err != nil {
		fmt.Println("RegisterValidation Error:", err)
		return nil
	}

	return c
}
