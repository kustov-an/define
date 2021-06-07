package main

import (
	"log"

	"github.com/kustov-an/define/common"
)

type Provider interface {
	Define(word string) (common.Definition, error)
}

type Printer interface {
	Print(definition common.Definition)
}

type Command struct {
	provider Provider
	printer  Printer
}

func (c *Command) run(word string) {
	defintion, err := c.provider.Define(word)
	if err != nil {
		log.Fatal(err)
	}
	c.printer.Print(defintion)
}
