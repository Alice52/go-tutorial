package validator

import (
	"fmt"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"testing"
)

// Validator tag: github.com/go-playground/validator/v10/baked_in.go
//
//	required: 字段不能为空
//
//	email || uri: 字段必须是有效的电子邮件地址
//
//	gt & gte || lt & lte 字段必须大于指定值
//
//	min || len || max: 字段长度必须等于指定长度
//	numeric || alpha || alphanum || hexadecimal: 字段必须是数字
//
//	ipv4 || ipv6: 字段必须是有效的IPv4地址

func TestStructTag(t *testing.T) {
	type Person struct {
		Name  string `json:"name" validate:"required,email"`  // gin 里是使用  binding 关键字
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"gte=18,lte=60"`
	}

	person := Person{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	}

	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	err = validate.Struct(person)
	if err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	fmt.Println("Valid person:", person)
}

func TestCustomTag(t *testing.T) {
	type Person struct {
		Name string `json:"name" validate:"CNumber"`
	}

	person := Person{
		Name: "234",
	}

	validate := validator.New()
	register(validate, "CNumber", numberTag)

	err := validate.Struct(person)
	if err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	fmt.Println("Valid person:", person)
}
