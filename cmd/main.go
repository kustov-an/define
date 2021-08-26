package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kustov-an/define/printer"
	"github.com/kustov-an/define/provider"
)

const usage = "Usage: define <word>"

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}
	verbose := flag.Bool("v", false, "extended output")
	flag.Parse()
	
	if len(flag.Args()) > 1 {
		fmt.Println(usage)
		return
	}
	word := flag.Arg(0)

	var prntr Printer
	if *verbose {
		prntr = &printer.ExtendedPrinter{}
	} else {
		prntr = &printer.SimplePrinter{}
	}

	command := &Command{
		provider: &provider.ApiProvider{
			Client: http.DefaultClient,
		},
		printer: prntr,
	}

	command.run(word)
}
