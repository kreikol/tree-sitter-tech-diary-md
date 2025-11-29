package tree_sitter_markdown_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_tech_diary "github.com/kreikol/tree-sitter-tech-diary-md/bindings/go"
)

func TestCanLoadBlockGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_tech_diary.Language())
	if language == nil {
		t.Errorf("Error loading Tech Diary block grammar")
	}
}

func TestCanLoadInlineGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_tech_diary.InlineLanguage())
	if language == nil {
		t.Errorf("Error loading Tech Diary inline grammar")
	}
}
