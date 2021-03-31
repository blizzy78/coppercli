package main

import (
	"context"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/blizzy78/copper/helpers"
	"github.com/blizzy78/copper/ranger"
	"github.com/blizzy78/copper/template"
)

type templateData struct {
	key   string
	value string
}

type dataFlags []templateData

type argError struct {
	msg string
}

func main() {
	var dataFl dataFlags
	flag.Var(&dataFl, "data", "template data (KEY=VALUE) (multiple)")

	flag.Parse()

	data := make(map[string]interface{}, len(dataFl))
	for _, d := range dataFl {
		data[d.key] = d.value
	}

	if err := render(os.Stdout, os.Stdin, data); err != nil {
		panic(err)
	}
}

func render(w io.Writer, r io.ReadCloser, data map[string]interface{}) error {
	l := template.LoaderFunc(func(name string) (io.ReadCloser, error) {
		return r, nil
	})

	rd := template.NewRenderer(l,
		template.WithScopeData("safe", helpers.Safe),
		template.WithScopeData("html", helpers.HTML),
		template.WithScopeData("has", helpers.Has),
		template.WithScopeData("len", helpers.Len),
		template.WithScopeData("hasPrefix", helpers.HasPrefix),
		template.WithScopeData("hasSuffix", helpers.HasSuffix),
		template.WithScopeData("range", ranger.New),
		template.WithScopeData("rangeFromTo", ranger.NewFromTo),
		template.WithScopeData("rangeInt", ranger.NewInt))

	return rd.Render(context.Background(), w, "template", data)
}

func (f *dataFlags) Set(v string) error {
	parts := strings.Split(v, "=")
	if len(parts) != 2 {
		return &argError{
			msg: "expected KEY=VALUE",
		}
	}

	key := strings.TrimSpace(parts[0])
	if key == "" {
		return &argError{
			msg: "expected KEY=VALUE",
		}
	}

	val := strings.TrimSpace(parts[1])

	*f = append(*f, templateData{
		key:   key,
		value: val,
	})
	return nil
}

func (f *dataFlags) String() string {
	return ""
}

func (e argError) Error() string {
	return e.msg
}
