package lib_test

import (
	"my-tools/thanksfulness/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckApiKey1(t *testing.T) {
	// given
	reqApiKey, thanksfulnessApiKey := "Abc123!", "Abc123!"
	// when
	actual, err := lib.CheckApiKey(reqApiKey, thanksfulnessApiKey)
	// then
	assert.Equal(t, true, actual)
	assert.Nil(t, err)
}

func TestCheckApiKey2(t *testing.T) {
	// given
	reqApiKey, thanksfulnessApiKey := "Abc123!", "Abc123"
	// when
	actual, err := lib.CheckApiKey(reqApiKey, thanksfulnessApiKey)
	// then
	assert.Equal(t, false, actual)
	assert.EqualError(t, err, "invalid api key")
}
