package pikchr

import (
	"bytes"
	"fmt"

	"github.com/gopikchr/gopikchr"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var toggleAttrName = []byte("toggle")
var limitWidthAttrName = []byte("limitwidth")

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
		toggle := false
		limitWidth := true

		var info []byte
		if n.Info != nil {
			info = n.Info.Segment.Value(src)
		}
		attrs := getAttributes(&n.FencedCodeBlock, info)
		if attrs != nil {
			if toggleAttr, ok := attrs.Get(toggleAttrName); ok {
				if val, ok := toggleAttr.(bool); ok {
					toggle = val
					if n.showToggleScript != nil {
						*n.showToggleScript = true
					}
				}
			}
			if limitWidthAttr, ok := attrs.Get(limitWidthAttrName); ok {
				if val, ok := limitWidthAttr.(bool); ok {
					limitWidth = val
				}
			}
		}
		fmt.Fprintf(w, "<div id='pikchr-%d' class='pikchr'", n.index)
		if toggle {
			fmt.Fprintf(w, " onclick=\"toggleHidden('pikchr-%d')\"", n.index)
		}
		w.WriteString(">\n")
		lines := n.Lines()
		var buf bytes.Buffer
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			buf.Write(line.Value(src))
		}

		zOut, width, _, _ := gopikchr.Convert(buf.String())
		if limitWidth {
			fmt.Fprintf(w, "<div style='max-width:%dpx'>\n%s</div>\n", width, zOut)
		} else {
			fmt.Fprintf(w, "<div>\n%s</div>\n", zOut)
		}
		if toggle {
			fmt.Fprintf(w, "<pre class='hidden'>\n")
			for i := 0; i < lines.Len(); i++ {
				line := lines.At(i)
				w.Write(line.Value(src))
			}
			fmt.Fprintf(w, "</pre>\n")
		}
		w.WriteString("</div>\n")
	}
	return ast.WalkContinue, nil
}

// pikchrJS is the javascript used to toggle viewing of Pikchr scripts with
// their source.
const pikchrJS = `<script>
  function toggleHidden(id){
    for(var c of document.getElementById(id).children){
      c.classList.toggle('hidden');
    }
  }
</script>`

// RenderScript renders pikchr.ScriptBlock nodes.
func (r *Renderer) RenderScript(w util.BufWriter, src []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	sb := node.(*ScriptBlock)
	if entering {
		if sb.showToggleScript {
			w.WriteString(pikchrJS)
		}
	}

	return ast.WalkContinue, nil
}
