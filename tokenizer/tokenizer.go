package tokenizer

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bbalet/stopwords"
	"github.com/kljensen/snowball"
)

const (
	english      = "english"
	englishShort = "en"
)

func Tokenize(inputString string) ([]string, error) {
	// remove stop words
	s := stopwords.CleanString(inputString, english, false)

	// lowercase
	s = strings.ToLower(s)

	// split the string into words
	sliceOfWords := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// stemming
	output := []string{}
	for _, w := range sliceOfWords {
		stemmed, err := snowball.Stem(w, english, true)
		if err != nil {
			return nil, fmt.Errorf("error stemming word: %s", w)
		}
		output = append(output, stemmed)
	}

	return output, nil
}
