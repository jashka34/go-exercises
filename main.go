package main

import (
	"fmt"

	"solution/solution11"
)

func main() {
	r := solution11.UserCreateRequest{}
	fmt.Println(solution11.Validate(r))
}
