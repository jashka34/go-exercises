package solution11

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	return ""
}
