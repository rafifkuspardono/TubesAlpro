package pencarian

import (
	"TubesAlpro/manajemenData"
)

func SeqSearch(T manajemenData.TabMakanan, n int, X string) int {
	var idx int = -1
	var i int
	for i = 0; i < n && idx == -1; i++ {
		if T[i].Nama == X {
			idx = i
		}
	}
	return idx
}

func BinarySearch(T manajemenData.TabMakanan, n int, X string) int {
	var left, right, mid int
	var found int = -1
	left = 0
	right = n - 1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if X < T[mid].Nama {
			right = mid - 1
		} else if X > T[mid].Nama {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}