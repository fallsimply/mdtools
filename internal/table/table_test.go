package table_test

import (
	"reflect"
	"testing"

	. "github.com/fallsimply/mdtools/internal/table"
)

var (
	tableStr = `|name|breed|eye color|
| ---- | :----: | ----: |
| maggie | basset hound | brown |
|flynn | shih tzu | black|`

	desired = Table{
		Headers:   []string{"name", "breed", "eye color"},
		Alignment: []Justify{Left, Center, Right},
		Rows: [][]string{
			{"maggie", "basset hound", "brown"},
			{"flynn", "shih tzu", "black"},
		},
		Width: []int{6, 12, 9},
	}
)

func TestParse(t *testing.T) {
	data := Parse(tableStr)
	if !reflect.DeepEqual(data, desired) {
		t.Errorf(`Result:\n%v\nDesired:\n%v\n`, data, desired)
	}
}

func TestJustify_String(t *testing.T) {
	if Left.String() != "Left" {
		t.Error(Left)
	}
	if Center.String() != "Center" {
		t.Error(Center)
	}
	if Right.String() != "Right" {
		t.Error(Right)
	}
}
