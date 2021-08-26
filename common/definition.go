package common

type Definition struct {
	Meanings []Meaning
}

type Meaning struct {
	PartOfSpeech string
	Meaning      string
	Synonyms     []string
}
