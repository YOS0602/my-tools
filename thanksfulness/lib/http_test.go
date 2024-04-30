package lib_test

import (
	"my-tools/thanksfulness/lib"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHeader(t *testing.T) {
	t.Parallel()
	// given
	current := http.Header{
		"X-Api-Version":   {"v1"},
		"Accept-Encoding": {"gzip"},
	}
	input := map[string][]string{
		"Authorization":   {"Bearer SECRET_TOKEN!!!"},
		"Accept-Language": {"ja", "en-US"},
		"Accept-Encoding": {"deflate"},
	}
	// when
	actual := lib.AddHeaders(current, input)
	// then
	assert.Equal(t, http.Header{
		"X-Api-Version":   {"v1"},
		"Accept-Encoding": {"gzip", "deflate"},
		"Authorization":   {"Bearer SECRET_TOKEN!!!"},
		"Accept-Language": {"ja", "en-US"},
	}, actual)
}
