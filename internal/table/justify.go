package table

const (
	// Left Align
	Left Justify = iota
	// Right Align
	Right
	// Center Align
	Center
)

// Justify is a number flag
type Justify int8

// Compute Justify flag from left and right booleans
func (ji *Justify) Compute(left, right bool) {
	switch {
	case right && left:
		*ji = Center
	case right:
		*ji = Right
	default:
		*ji = Left
	}
}

func (ji Justify) String() string {
	switch ji {
	case Right:
		return "Right"
	case Center:
		return "Center"
	default:
		return "Left"
	}
}

func parseJustifyTokens(tokens []string) (justify []Justify) {
	for _, token := range tokens {
		var start, left, right bool
		for _, char := range token {
			switch char {
			case ':':
				if !start {
					start = true
					left = true
				} else {
					right = true
				}
			case '-':
				if !start {
					start = true
				}
			}
		}
		var item Justify
		item.Compute(left, right)
		justify = append(justify, item)
	}
	return
}
