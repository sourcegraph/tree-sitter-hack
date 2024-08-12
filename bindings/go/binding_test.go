package tree_sitter_hack_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-hack"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_hack.Language())
	if language == nil {
		t.Errorf("Error loading Hack grammar")
	}
}
