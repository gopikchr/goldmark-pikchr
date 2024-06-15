package pikchr

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func TestRenderer_Block(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		give string
		want string
	}{
		{
			desc: "empty",
			give: "",
			want: "<div id='pikchr-0' class='pikchr'>\n<div class='pikchr-svg' style='max-width:0px'>\n<!-- empty pikchr diagram -->\n</div>\n</div>\n",
		},
		{
			desc: "graph",
			give: `box "Test"`,
			want: trimLeadingSpace(`<div id='pikchr-0' class='pikchr'>
				<div class='pikchr-svg' style='max-width:112px'>
				<svg xmlns='http://www.w3.org/2000/svg' viewBox="0 0 112.32 76.32">
				<path d="M2,74L110,74L110,2L2,2Z"  style="fill:none;stroke-width:2.16;stroke:rgb(0,0,0);" />
				<text x="56" y="38" text-anchor="middle" fill="rgb(0,0,0)" dominant-baseline="central">Test</text>
				</svg>
				</div>
				</div>
				`),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			r := renderer.NewRenderer(
				renderer.WithNodeRenderers(
					util.Prioritized(&Renderer{}, 100),
				),
			)

			reader := text.NewReader([]byte(tt.give))
			segs := text.NewSegments()
			for {
				line, seg := reader.PeekLine()
				if line == nil {
					break
				}

				segs.Append(seg)
				reader.AdvanceLine()
			}
			give := new(Block)
			give.SetLines(segs)

			var buff bytes.Buffer
			w := bufio.NewWriter(&buff)

			assert.NoError(t, r.Render(w, reader.Source(), give), "Render")
			assert.Equal(t, tt.want, buff.String())
		})
	}
}

func TestRenderer_Script_Toggle(t *testing.T) {
	want := `<script>
  function toggleHidden(id){
    for(var c of document.getElementById(id).children){
      c.classList.toggle('hidden');
    }
  }
</script>`
	r := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{}, 100),
		),
	)

	var buff bytes.Buffer
	w := bufio.NewWriter(&buff)

	assert.NoError(t,
		r.Render(w, nil /* src */, &ScriptBlock{showToggleScript: true}))
	assert.Equal(t, want, buff.String())
}

func TestRenderer_Script_No_Toggle(t *testing.T) {
	want := ""
	r := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(&Renderer{}, 100),
		),
	)

	var buff bytes.Buffer
	w := bufio.NewWriter(&buff)

	assert.NoError(t,
		r.Render(w, nil /* src */, &ScriptBlock{}))
	assert.Equal(t, want, buff.String())
}
