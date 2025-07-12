package solution24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatinLetters(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", latinLetters(" abc1"))
	a.Equal("", latinLetters(" привет"))
	a.Equal("test", latinLetters("11 ! t e    s t %)"))
	a.Equal("John", latinLetters("John Уолтер"))
	a.Equal("word", latinLetters("wo[r]d"))
}
