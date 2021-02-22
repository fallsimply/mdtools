package table

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var (
	tblStr = `|name|breed|eye color|
| ---- | :----: | ----: |
| maggie | basset hound | brown |
|flynn | shih tzu | black|`

	tblHeaders = []string{"name", "breed", "eye color"}
	tblAlign   = []Justify{Left, Center, Right}
	tblWidth   = []int{6, 12, 9}
	tblRows    = [][]string{
		{"maggie", "basset hound", "brown"},
		{"flynn", "shih tzu", "black"},
	}

	desired = Table{
		Headers:   tblHeaders,
		Alignment: tblAlign,
		Rows:      tblRows,
		Width:     tblWidth,
	}
)

func TestParse(t *testing.T) {
	data := Parse(tblStr)
	if !reflect.DeepEqual(data, desired) {
		t.Errorf(`Result:\n`+`%v\n`+`Desired:\n`+`%v\n`, data, desired)
	}
}

func TestTokenizeLine(t *testing.T) {
	var width []int
	line := `| maggie | basset hound | brown |`
	if tokens := tokenizeLine(line, &width); !reflect.DeepEqual(tokens, tblRows[0]) {
		t.Errorf("%q: %s", line, tokens)
	}
}

func TestParseJustifyTokens(t *testing.T) {
	cases := []struct {
		str    string
		Result Justify
	}{
		{str: "---", Result: Left},
		{str: ":---", Result: Left},
		{str: ":---:", Result: Center},
		{str: "-----:", Result: Right},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if token := parseJustifyTokens([]string{tc.str})[0]; token != tc.Result {
				t.Errorf("%q: %v", tc.str, token)
			}
		})
	}
}

func TestJustify_String(t *testing.T) {
	cases := []struct {
		Item   Justify
		Result string
	}{
		{Item: Left, Result: "Left"},
		{Item: Right, Result: "Right"},
		{Item: Center, Result: "Center"},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tc.Item.String() != tc.Result {
				t.Errorf("(%s, %[1]d): %s", tc.Item, tc.Result)
			}
		})
	}
}

func TestJustify_Compute(t *testing.T) {
	var j Justify
	cases := []struct {
		Left, Right bool
		Result      Justify
	}{
		{Left: false, Right: false, Result: Left},
		{Left: false, Right: false, Result: Left},
		{Left: true, Right: true, Result: Center},
		{Left: false, Right: true, Result: Right},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			if j.Compute(tc.Left, tc.Right); j != tc.Result {
				t.Errorf("%s:\n"+"Left: %t\n"+"Right: %t\n"+"Align: %v\n", t.Name(), tc.Left, tc.Right, j)
			}
		})
	}
}
