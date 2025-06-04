package pengurutan

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func InsertionSortStok(A *manajemenData.TabMakanan, N int) {
	var pass, i int
	var temp manajemenData.BahanMakanan

	pass = 1
	for pass < N {
		i = pass
		temp = A[pass]
		for i > 0 && temp.JumlahStok < A[i-1].JumlahStok {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
	fmt.Println("Data diurutkan berdasarkan jumlah stok (terkecil ke terbesar).")
}
