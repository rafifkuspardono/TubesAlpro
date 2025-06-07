package pengurutan

import (
	"TubesAlpro/manajemenData"
)

func SelectionNama(A *manajemenData.TabMakanan, N int) {
	var pass, i, idx int
	var temp manajemenData.BahanMakanan

	pass = 0
	for pass < N-1 {
		idx = pass
		i = pass + 1

		for i < N {
			if A[i].Nama < A[idx].Nama {
				idx = i
			}
			i = i + 1
		}

		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}
