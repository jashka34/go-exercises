package solution30

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	var cur CreateUserRequest
	err := json.Unmarshal(requestBody, &cur)
	fmt.Printf("cur: %+v\n", cur)
	if cur.Email == "" {
		err = errEmailRequired
	} else if cur.Password == "" {
		err = errPasswordRequired
	} else if cur.Password != cur.PasswordConfirmation {
		err = errPasswordDoesNotMatch
	}
	fmt.Println("err:", err)
	if err != nil {
		cur = CreateUserRequest{}
	}
	return cur, err
}
