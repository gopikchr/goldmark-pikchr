package pikchr

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Extender adds support for Pikchr diagrams to a Goldmark Markdown parser.
//
// Use it by installing it to the goldmark.Markdown object upon creation.
//
//	goldmark.New(
//	  // ...
//	  goldmark.WithExtensions(
//	    // ...
//	    &pikchr.Exender{},
//	  ),
//	)
type Extender struct {
	ToggleDefault     bool // If true, turn toggling on by default
	LimitWidthDefault bool // If true, turn limitwidth on by default
}

// Extend extends the provided Goldmark parser with support for Pikchr
// diagrams.
func (e *Extender) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(&Transformer{}, 1000),
		),
	)
	md.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(
				&Renderer{ToggleDefault: e.ToggleDefault, LimitWidthDefault: e.LimitWidthDefault}, 100),
		),
	)
}
