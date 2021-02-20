package html

import (
	"strings"

	core "golang.org/x/net/html"
)

// Render turns an element into an html string
func Render(tree *core.Node) string {
	var b strings.Builder
	core.Render(&b, tree)
	return b.String()
}
