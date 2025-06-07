package laporan

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func LaporanTotal(A manajemenData.TabMakanan, n int) {
	var total, digunakan, tersedia int
	var i int

	total = 0
	digunakan = 0
	tersedia = 0

	for i = 0; i < n; i++ {
		total++
		if A[i].SudahDipakai {
			digunakan++
		} else {
			tersedia++
		}
	}

	fmt.Println("\n--- LAPORAN STOK BAHAN MAKANAN ---")
	fmt.Printf("Total bahan        : %d\n", total)
	fmt.Printf("Sudah dipakai      : %d\n", digunakan)
	fmt.Printf("Belum dipakai      : %d\n", tersedia)
}
