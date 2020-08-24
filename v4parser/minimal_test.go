package v4parser_test

import (
	"testing"

	"github.com/jlj5aj/antlr4-issue/v4parser"
)

func BenchmarkSlowQuery(b *testing.B) {
	v4parser.WalkTree(`keyword: { I have often thought that nothing would do more extensive good at small expense than the establishment of a small circulating library in every county, to consist of a few well-chosen books, to be lent to the people of the country under regulations as would secure their safe return in due time. }`)
}
