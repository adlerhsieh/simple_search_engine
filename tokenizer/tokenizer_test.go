package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "I have a book",
			output: []string{"i", "book"},
		},
		{
			input:  "The code is organized as follows",
			output: []string{"code", "organ", "follow"},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			output, err := Tokenize(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.output, output)
		})
	}
}
