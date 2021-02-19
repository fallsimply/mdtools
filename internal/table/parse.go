package table

import (
	"strings"
)

// Table is a structured data representation of a markdown table
type Table struct {
	Headers   []string
	Alignment []Justify
	Rows      [][]string
	Width     []int
}

// Parse turns a markdown table string into structured data
func Parse(in string) (tbl Table) {
	for i, line := range strings.Split(in, "\n") {
		switch i {
		case 0:
			tbl.Headers = tokenizeLine(line, &tbl.Width)
		case 1:
			tbl.Alignment = parseJustifyTokens(tokenizeLine(line, &tbl.Width))
		default:
			tbl.Rows = append(tbl.Rows, tokenizeLine(line, &tbl.Width))
		}
	}
	return
}

func tokenizeLine(line string, width *[]int) (tokens []string) {
	line = strings.Trim(line, "|")
	tokens = strings.Split(line, "|")

	if *width == nil {
		*width = make([]int, len(tokens))
	}

	for i, val := range tokens {
		tokens[i] = strings.TrimSpace(val)
		val = tokens[i]
		if len(val) > (*width)[i] {
			(*width)[i] = len(val)
		}
	}

	return
}
