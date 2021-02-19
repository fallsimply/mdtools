package main

import (
	"fmt"

	tbl "github.com/fallsimply/mdtools/internal/table"
)

func main() {
	table := tbl.Parse(`||breed|eye color|
| ---- | :----: | ----: |
| maggie | basset hound | brown |
|flynn | shih tzu | black|`)
	fmt.Printf("%q\n%v\n%v\n%q\n", table.Headers, table.Alignment, table.Width, table.Rows)
}
