package html_test

import (
	"strings"
	"testing"

	"github.com/fallsimply/mdtools/internal/html"
	stdhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestNewElement(t *testing.T) {
	cases := []struct {
		Elem   string
		Attr   map[string]string
		Result string
	}{
		{Elem: "p", Result: `<p></p>`},
		{Elem: "a", Attr: map[string]string{"href": "#"}, Result: `<a href="#"></a>`},
	}
	for _, tc := range cases {
		t.Run(`Tag_`+tc.Elem, func(t *testing.T) {
			elem := html.NewElement(tc.Elem, tc.Attr)
			if txt := html.Render(elem); txt != tc.Result {
				t.Fatal(txt)
			}
		})
	}
}

func TestNewText(t *testing.T) {
	elem := html.NewText("test")
	if txt := html.Render(elem); txt != "test" {
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
	doc, err := stdhtml.ParseFragment(strings.NewReader(`<p>text</p>`), &stdhtml.Node{Type: stdhtml.ElementNode, Data: "body", DataAtom: atom.Body})
	if err != nil {
		t.Fatal(err)
	}
	if txt := html.Render(doc[0]); txt != `<p>text</p>` {
		t.Fatal(txt)
	}
}
