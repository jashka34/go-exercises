package solution23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isASCII(" abc1"))
	a.Equal(false, isASCII("хай"))
	a.Equal(false, isASCII("hello \U0001F970"))
}
