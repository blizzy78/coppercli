package main

import (
	"context"
	"errors"
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

func render(w io.Writer, r io.Reader, data map[string]interface{}) error {
	l := template.LoaderFunc(func(name string) (io.Reader, error) {
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
		return errors.New("expected KEY=VALUE")
	}

	key := strings.TrimSpace(parts[0])
	if len(key) <= 0 {
		return errors.New("expected KEY=VALUE")
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
