package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kustov-an/define/printer"
	"github.com/kustov-an/define/provider"
)

const usage = "Usage: define <word>"

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}
	word := os.Args[1]

	command := &Command{
		provider: &provider.ApiProvider{
			Client: http.DefaultClient,
		},
		printer: &printer.SimplePrinter{},
	}

	command.run(word)
}
