package tuple

type tuple struct{ elements []any }

func (t tuple) GetN(n int) (any, bool) {
	if n >= 1 && n <= len(t.elements) {
		return t.elements[n-1], true
	}
	return nil, false
}
