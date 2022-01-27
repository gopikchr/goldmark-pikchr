package pikchr

import (
	"bytes"

	"github.com/gopikchr/gopikchr"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer renders Pikchr diagrams as HTML/SVG.
type Renderer struct {
}

// RegisterFuncs registers the renderer for Pikchr blocks with the provided
// Goldmark Registerer.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.Render)
	reg.Register(ScriptKind, r.RenderScript)
}

// Render renders pikchr.Block nodes.
func (*Renderer) Render(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*Block)
	if entering {
		// w.WriteString(`<div class="pikchr">`)
		lines := n.Lines()
		var buf bytes.Buffer
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			buf.Write(line.Value(src))
		}

		zOut, _, _, _ := gopikchr.Convert(buf.String())
		w.WriteString(zOut)
	}
	return ast.WalkContinue, nil
}

// pikchrJS is the javascript used to toggle viewing of Pikchr scripts with
// their source.
const pikchrJS = `<script>
</script>`

// RenderScript renders pikchr.ScriptBlock nodes.
func (r *Renderer) RenderScript(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = node.(*ScriptBlock) // sanity check
	if entering {
		w.WriteString(pikchrJS)
	}

	return ast.WalkContinue, nil
}
