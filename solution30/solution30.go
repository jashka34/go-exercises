package solution30

import "errors"

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
	return cur, nil
}
