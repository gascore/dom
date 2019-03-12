package dom

import (
	"github.com/gascore/dom/js"
)

func AsTextNode(v js.Value) *TextNode {
	if !v.Valid() {
		return nil
	}
	return &TextNode{NodeBase{v: v}}
}

type TextNode struct {
	NodeBase
}