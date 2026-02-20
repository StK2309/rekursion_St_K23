package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	// TODO
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	if len(s) >= len(seq) && s[:len(seq)] == seq {
		return true
	}
	if len(s) > 1 {
		return Contains(s[1:], seq)
	}
	return false
}
