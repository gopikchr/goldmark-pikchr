package pikchr

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Transformer transforms a Goldmark Markdown AST with support for Pikchr
// diagrams. It makes the following transformations:
//
//  - replace pikchr code blocks with pikchr.Block nodes
//  - add a pikchr.ScriptBlock node if the document uses Pikchr
type Transformer struct {
}

var _pikchr = []byte("pikchr")

// Transform transforms the provided Markdown AST.
func (*Transformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {
	var (
		hasScript    bool
		pikchrBlocks []*ast.FencedCodeBlock
	)

	// Collect all blocks to be replaced without modifying the tree.
	ast.Walk(doc, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
		if !enter {
			return ast.WalkContinue, nil
		}

		// For multiple transforms.
		if _, ok := node.(*ScriptBlock); ok {
			hasScript = true
			return ast.WalkContinue, nil
		}

		cb, ok := node.(*ast.FencedCodeBlock)
		if !ok {
			return ast.WalkContinue, nil
		}

		lang := cb.Language(reader.Source())
		if !bytes.Equal(lang, _pikchr) {
			return ast.WalkContinue, nil
		}

		pikchrBlocks = append(pikchrBlocks, cb)
		return ast.WalkContinue, nil
	})

	// Nothing to do.
	if len(pikchrBlocks) == 0 {
		return
	}

	for _, cb := range pikchrBlocks {
		b := new(Block)
		b.SetLines(cb.Lines())

		parent := cb.Parent()
		if parent != nil {
			parent.ReplaceChild(parent, cb, b)
		}
	}

	if !hasScript && false { // no scripts for now
		doc.AppendChild(doc, &ScriptBlock{})
	}
}
