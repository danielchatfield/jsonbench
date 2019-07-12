package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path"
	"text/template"
)

const NumFields = 10000

func main() {
	templatePath := path.Join(".", "gen", "jsonbench.go.tmpl")
	t := template.New(path.Base(templatePath))

	t = t.Funcs(template.FuncMap{
		"Fields": func() []string {
			items := make([]string, NumFields)
			for i := 0; i < NumFields; i++ {
				items[i] = fmt.Sprintf("A%.4d", i)
			}
			return items
		},
	})

	t, err := t.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, nil); err != nil {
		log.Fatal(err)
	}

	bts, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("gofmt: %v", err)
	}

	err = ioutil.WriteFile("jsonbench.go", bts, 0644)
	if err != nil {
		log.Fatalf("Error writing: %v", err)
	}
}