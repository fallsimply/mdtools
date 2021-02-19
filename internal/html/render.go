package html

import (
	"strings"

	"golang.org/x/net/html"
)

// Render turns a parsed dom tree into an html string
func Render(tree *html.Node) string {
	var b strings.Builder
	html.Render(&b, tree)
	/*
	 * removes html, head, and body tags
	 *    len(`<html><head></head><body>`) == 25
	 *    len(`</body></html>`) == 14
	 */
	return string([]byte(b.String())[25 : b.Len()-14])
}
