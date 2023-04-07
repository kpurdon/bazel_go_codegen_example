package main

import (
	"encoding/json"
	"flag"
	"os"
	"text/template"
)

var code = `
package hello

import "fmt"

func Hello(name string) {
        fmt.Printf("Hello, %s{{ .Punctuation }}\n", name)
}
`

type Config struct {
	Punctuation string `json:"punctuation"`
}

func main() {
	configFile := flag.String("config-file", "gen/config.json", "config.json file")
	outputFile := flag.String("output-file", "hello.go", "output go file")
	flag.Parse()

	cf, err := os.Open(*configFile)
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	var cfg Config
	err = json.NewDecoder(cf).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	t, err := template.New("code").Parse(code)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(*outputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, &cfg)
	if err != nil {
		panic(err)
	}
}
