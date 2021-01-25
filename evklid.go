package evklid

// Evklid вычисляет наибольший общий делитель, используя алгоритм Евклида.
func Evklid(i int, j int) int {
	for i != j {
		if i > j {
			i -= j
		} else {
			j -= i
		}
	}
	return i
}
