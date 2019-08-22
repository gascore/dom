package dom

import "github.com/gascore/dom/js"

func GetDocument() *Document {
	doc := js.Get("document")
	if !doc.Valid() {
		return nil
	}
	return &Document{NodeBase{v: doc}}
}

func NewElement(tag string) *Element {
	return Doc.CreateElement(tag)
}

func NewElementNS(ns, tag string) *Element {
	return Doc.CreateElementNS(ns, tag)
}

var _ Node = (*Document)(nil)

type Document struct {
	NodeBase
}

func (d *Document) CreateElement(tag string) *Element {
	v := d.v.Call("createElement", tag)
	return AsElement(v)
}
func (d *Document) CreateTextNode(value string) *TextNode {
	v := d.v.Call("createTextNode", value)
	return AsTextNode(v)
}
func (d *Document) CreateElementNS(ns string, tag string) *Element {
	v := d.v.Call("createElementNS", ns, tag)
	return AsElement(v)
}
func (d *Document) GetElementById(id string) *Element {
	v := d.v.Call("getElementById", id)
	return AsElement(v)
}
func (d *Document) GetElementsByTagName(tag string) NodeList {
	v := d.v.Call("getElementsByTagName", tag)
	return AsNodeList(v)
}
func (d *Document) QuerySelector(qu string) *Element {
	v := d.v.Call("querySelector", qu)
	return AsElement(v)
}
func (d *Document) QuerySelectorAll(qu string) NodeList {
	v := d.v.Call("querySelectorAll", qu)
	return AsNodeList(v)
}
