package html

import (
	"strings"

	stdhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// NewElement creates a new element from a name and attribute map
func NewElement(name string, attributes map[string]string) (elem *stdhtml.Node) {
	elem = &stdhtml.Node{
		Type:     stdhtml.ElementNode,
		Data:     name,
		DataAtom: atom.Lookup([]byte(strings.Title(name))),
		Attr:     make([]stdhtml.Attribute, 0),
	}

	for name, value := range attributes {
		elem.Attr = append(elem.Attr, stdhtml.Attribute{Key: name, Val: value})
	}

	return
}

// NewText creates a new text node
func NewText(text string) *stdhtml.Node {
	return &stdhtml.Node{
		Type: stdhtml.TextNode,
		Data: text,
	}
}
