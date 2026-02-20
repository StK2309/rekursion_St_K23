package strings

import "strings"

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// TODO
	if strings.HasPrefix(s, seq) {
		return true
	}
	return false
}
