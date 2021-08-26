package printer

import (
	"fmt"

	"github.com/kustov-an/define/common"
)

type SimplePrinter struct {
}

func (s *SimplePrinter) Print(definition common.Definition) {
	var i uint = 1
	for _, meaning := range definition.Meanings {
		fmt.Printf("%v. %v\n", i, meaning.Meaning)
		i++
	}
}
