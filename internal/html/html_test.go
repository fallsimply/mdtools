package html_test

import (
	"strings"
	"testing"

	"github.com/fallsimply/mdtools/internal/html"
	stdhtml "golang.org/x/net/html"
)

func TestNewElement_paragraph(t *testing.T) {
	elem := html.NewElement("p", nil)
	if txt := html.Render(elem); txt != `<p></p>` {
		t.Fatal(txt)
	}
}

func TestNewElement_link(t *testing.T) {
	elem := html.NewElement("a", map[string]string{
		"href": "#",
	})
	if txt := html.Render(elem); txt != `<a href="#"></a>` {
		t.Fatal(txt)
	}
}

func TestNewText(t *testing.T) {
	elem := html.NewText("test")
	if txt := html.Render(elem); txt != `test` {
		t.Fatal(txt)
	}
}

func Test_elementWithText(t *testing.T) {
	elem := html.NewElement("p", nil)
	elem.AppendChild(html.NewText("text"))

	if txt := html.Render(elem); txt != `<p>text</p>` {
		t.Fatal(txt)
	}
}

func TestRender(t *testing.T) {
	doc, err := stdhtml.Parse(strings.NewReader(`<p>text</p>`))
	if err != nil {
		t.Fatal(err)
	}
	if txt := html.Render(doc); txt != `<html><head></head><body><p>text</p></body></html>` {
		t.Fatal(txt)
	}
}
