package pengurutan

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func LebihAwal(a, b manajemenData.BahanMakanan) bool {
	if a.TahunKadaluwarsa < b.TahunKadaluwarsa {
		return true
	} else if a.TahunKadaluwarsa == b.TahunKadaluwarsa {
		if a.BulanKadaluwarsa < b.BulanKadaluwarsa {
			return true
		} else if a.BulanKadaluwarsa == b.BulanKadaluwarsa {
			return a.TanggalKadaluwarsa < b.TanggalKadaluwarsa
		}
	}
	return false
}

func SelectionSortKadaluarsa(A *manajemenData.TabMakanan, n int) {
	var pass, idx, i int
	var temp manajemenData.BahanMakanan

	pass = 0
	for pass < n-1 {
		idx = pass
		i = pass + 1
		for i < n {
			if LebihAwal(A[i], A[idx]) {
				idx = i
			}
			i++
		}
		temp = A[pass]
		A[pass] = A[idx]
		A[idx] = temp
		pass++
	}
	fmt.Println("Data diurutkan berdasarkan tanggal kadaluarsa (paling awal ke paling akhir).")
}
