package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	// TODO
	if pos < 0 {
		return list
	}
	if pos == 0 {
		return list[1:]
	}
	if Empty(list) {
		return list
	}
	return append(list[:1], RemoveElement(list[1:], pos-1)...)
}
