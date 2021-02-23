package html

import (
	"strings"

	stdhtml "golang.org/x/net/html"
)

// Render turns an element tree into an html string
func Render(tree *stdhtml.Node) string {
	var b strings.Builder
	stdhtml.Render(&b, tree)
	return b.String()
}
