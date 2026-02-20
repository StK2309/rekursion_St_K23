package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	// TODO
	if len(list) == 0 {
		return -1
	}
	mid := len(list) / 2
	if list[mid] == x {
		return mid
	}
	if list[mid] > x {
		return FindSorted(list[:mid], x)
	}
	return mid + 1 + FindSorted(list[mid+1:], x)
}
