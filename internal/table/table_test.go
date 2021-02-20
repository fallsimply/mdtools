package table_test

import (
	"reflect"
	"testing"

	tbl "github.com/fallsimply/mdtools/internal/table"
)

var (
	tableStr = `|name|breed|eye color|
| ---- | :----: | ----: |
| maggie | basset hound | brown |
|flynn | shih tzu | black|`

	desired = tbl.Table{
		Headers:   []string{"name", "breed", "eye color"},
		Alignment: []tbl.Justify{tbl.Left, tbl.Center, tbl.Right},
		Rows: [][]string{
			{"maggie", "basset hound", "brown"},
			{"flynn", "shih tzu", "black"},
		},
		Width: []int{6, 12, 9},
	}
)

func TestParse(t *testing.T) {
	data := tbl.Parse(tableStr)
	if !reflect.DeepEqual(data, desired) {
		t.Errorf(`Result:\n%v\nDesired:\n%v\n`, data, desired)
	}
}

func TestJustify_String(t *testing.T) {
	if tbl.Left.String() != "Left" {
		t.Errorf("%s: %v", "Center", tbl.Left.String())
	}
	if tbl.Center.String() != "Center" {
		t.Errorf("%s: %v", "Center", tbl.Center.String())
	}
	if tbl.Right.String() != "Right" {
		t.Errorf("%s: %v", "Center", tbl.Right.String())
	}
}
