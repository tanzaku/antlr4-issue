package v4parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func WalkTree(query string) {
	input := antlr.NewInputStream(query)
	tokens := antlr.NewCommonTokenStream(NewVirgoQueryLexer(input), 0)
	tree := NewVirgoQuery(tokens)
	antlr.NewParseTreeWalker().Walk(new(BaseVirgoQueryListener), tree.Query())
}
