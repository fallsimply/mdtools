package html

import (
	"strings"

	core "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// NewElement creates a new element from the element name and attribute map
func NewElement(name string, attributes map[string]string) (elem *core.Node) {
	elem = new(core.Node)
	elem.Data = name
	elem.DataAtom = atom.Lookup([]byte(strings.Title(name)))
	elem.Type = core.ElementNode
	elem.Attr = make([]core.Attribute, 0)

	for name, value := range attributes {
		elem.Attr = append(elem.Attr, core.Attribute{
			Key: name,
			Val: value,
		})
	}

	return
}

// NewText creates a new text node from the text
func NewText(text string) (elem *core.Node) {
	elem = new(core.Node)
	elem.Data = text
	elem.Type = core.TextNode

	return
}
