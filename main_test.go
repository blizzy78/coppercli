package main

import (
	"io"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)

	tmpl := `<% safe(len(s)) %>`
	r := io.NopCloser(strings.NewReader(tmpl))
	buf := strings.Builder{}

	err := render(&buf, r, map[string]interface{}{
		"s": "hello",
	})

	is.NoErr(err)
	is.Equal("5", buf.String())
}
