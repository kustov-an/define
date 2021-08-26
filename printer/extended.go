package printer

import (
	"fmt"
	"strings"

	"github.com/kustov-an/define/common"
)

type ExtendedPrinter struct {
}

func (s *ExtendedPrinter) Print(definition common.Definition) {
	var i uint = 1
	for _, meaning := range definition.Meanings {
		fmt.Printf("%v. [%v] %v\n", i, meaning.PartOfSpeech, meaning.Meaning)
		if len(meaning.Synonyms) > 0 {
			fmt.Printf("Synonyms: %v\n", strings.Join(meaning.Synonyms, ", "))
		}
		i++
	}
}
