package sample_test

import (
	"testing"

	"my-tools/thanksfulness/sample"

	"github.com/stretchr/testify/assert"
)

func TestCry(t *testing.T) {
	c1 := sample.NewCat("ニャーニャー")
	assert.Equal(t, "ニャーニャー", c1.Cry())
}
