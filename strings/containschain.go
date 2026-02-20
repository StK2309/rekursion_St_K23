package strings

import "strings"

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {

	if strings.HasPrefix(s, strings.Repeat(symbol, count)) {
		return true
	}
	if len(s) < len(symbol) {
		return false
	}
	return ContainsChain(s[1:], symbol, count)
}
