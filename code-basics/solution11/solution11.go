package solution11

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

type UserCreateRequest1 struct {
	FirstName string `validate:"required,excludesall= "`
	Age       int    `validate:"gt=0,lte=150"`
}

func Validate(req UserCreateRequest) string {
	if strings.Contains(req.FirstName, " ") || req.FirstName == "" {
		return "invalid request"
	}
	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}
	return ""
}

func Validate1(req UserCreateRequest1) string {
	v := validator.New()
	err := v.Struct(req)
	// fmt.Println(err)
	if err != nil {
		return "invalid request"
	} else {
		return ""
	}
}
