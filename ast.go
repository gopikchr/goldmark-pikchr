package pikchr

import "github.com/yuin/goldmark/ast"

// Kind is the kind of a Pikchr block.
var Kind = ast.NewNodeKind("PikchrBlock")

// Block is a Pikchr block.
//
//  ```pikchr
//  box color red wid 2.6in \
//      "Click on any diagram on this page" big \
//      "to see the Pikchr source text" big
//  ```
//
// Its raw contents are the plain text of the Pikchr diagram.
type Block struct {
	ast.FencedCodeBlock
	index            int
	showToggleScript *bool
}

// IsRaw reports that this block should be rendered as-is.
func (*Block) IsRaw() bool { return true }

// Kind reports that this is a PikchrBlock.
func (*Block) Kind() ast.NodeKind { return Kind }

// Dump dumps the contents of this block to stdout.
func (b *Block) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}

// ScriptKind is the kind of a Pikchr Script block.
var ScriptKind = ast.NewNodeKind("PikchrScriptBlock")

// ScriptBlock marks where the Pikchr Javascript will be included.
//
// This is a placeholder and does not contain anything.
type ScriptBlock struct {
	ast.BaseBlock
	showToggleScript bool
}

// IsRaw reports that this block should be rendered as-is.
func (*ScriptBlock) IsRaw() bool { return true }

// Kind reports that this is a PikchrScriptBlock.
func (*ScriptBlock) Kind() ast.NodeKind { return ScriptKind }

// Dump dumps the contents of this block to stdout.
func (b *ScriptBlock) Dump(src []byte, level int) {
	ast.DumpHelper(b, src, level, nil, nil)
}
