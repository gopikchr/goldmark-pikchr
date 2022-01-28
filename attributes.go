package pikchr

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Almost everything in this file is copied from
// https://github.com/yuin/goldmark-highlighting

// ImmutableAttributes is a read-only interface for ast.Attributes.
type ImmutableAttributes interface {
	// Get returns (value, true) if an attribute associated with given
	// name exists, otherwise (nil, false)
	Get(name []byte) (interface{}, bool)

	// GetString returns (value, true) if an attribute associated with given
	// name exists, otherwise (nil, false)
	GetString(name string) (interface{}, bool)

	// All returns all attributes.
	All() []ast.Attribute
}

type immutableAttributes struct {
	n ast.Node
}

func (a *immutableAttributes) Get(name []byte) (interface{}, bool) {
	return a.n.Attribute(name)
}

func (a *immutableAttributes) GetString(name string) (interface{}, bool) {
	return a.n.AttributeString(name)
}

func (a *immutableAttributes) All() []ast.Attribute {
	if a.n.Attributes() == nil {
		return []ast.Attribute{}
	}
	return a.n.Attributes()
}

func getAttributes(node *ast.FencedCodeBlock, infostr []byte) ImmutableAttributes {
	if node.Attributes() != nil {
		return &immutableAttributes{node}
	}
	if infostr != nil {
		attrStartIdx := -1

		for idx, char := range infostr {
			if char == '{' {
				attrStartIdx = idx
				break
			}
		}
		if attrStartIdx > 0 {
			n := ast.NewTextBlock() // dummy node for storing attributes
			attrStr := infostr[attrStartIdx:]
			if attrs, hasAttr := parser.ParseAttributes(text.NewReader(attrStr)); hasAttr {
				for _, attr := range attrs {
					n.SetAttribute(attr.Name, attr.Value)
				}
				return &immutableAttributes{n}
			}
		}
	}
	return nil
}
