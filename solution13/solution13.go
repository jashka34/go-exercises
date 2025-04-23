package solution13

import "fmt"

type Voicer interface {
	Voice() string
}

type Cat struct {
	// …
}

func (c *Cat) Voice() string {
	return "Мяу"
}

type Cow struct {
	// …
}

func (c *Cow) Voice() string {
	return "Мууу"
}

type Dog struct {
	// …
}

func (c *Dog) Voice() string {
	return "Гав"
}

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}
	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}
