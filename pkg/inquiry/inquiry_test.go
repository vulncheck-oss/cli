package inquiry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterASCII(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard ASCII", "Hello, World!", "Hello, World!"},
		{"With spaces", "John Doe's Computer", "John Doe's Computer"},
		{"Left single quote", "John\u2018s Computer", "John's Computer"},
		{"Right single quote", "John\u2019s Computer", "John's Computer"},
		{"Single high-reversed-9 quote", "John\u201Bs Computer", "John's Computer"},
		{"Grave accent", "John\u0060s Computer", "John's Computer"},
		{"Acute accent", "John\u00B4s Computer", "John's Computer"},
		{"Left double quote", "\u201CJohn's Computer\u201D", "\"John's Computer\""},
		{"Right double quote", "\u201DJohn's Computer\u201C", "\"John's Computer\""},
		{"Double high-reversed-9 quote", "\u201FJohn's Computer\u201F", "\"John's Computer\""},
		{"Mixed quotes", "\u2018John\u2019s \u201CComputer\u201D", "'John's \"Computer\""},
		{"Non-ASCII characters", "John's Café Latté", "John's Caf Latt"},
		{"Emoji", "John's 👨‍💻 Computer", "John's  Computer"},
		{"Mixed ASCII and non-ASCII", "John's Æsthetic Møødŷ Cømputer", "John's sthetic Md Cmputer"},
		{"Empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterASCII(tt.input)
			assert.Equal(t, tt.expected, result, "Input: %q", tt.input)
		})
	}
}
