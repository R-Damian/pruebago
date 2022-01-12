package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {

	result := ExampleTest()

	assert.Equal(t, "Ok", result)
}
