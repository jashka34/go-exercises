package solution13

import "fmt"

type Voicer interface {
	Voice() string
}

type Cat struct {
	// …
}

type Cow struct {
	// …
}

type Dog struct {
	// …
}

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}
	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}
