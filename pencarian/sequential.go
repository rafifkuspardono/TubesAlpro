package pencarian

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func SequentialSearch(A manajemenData.TabMakanan, n int, X string) int {
	var idx int = -1
	var i int
	for i = 0; i < n && idx == -1; i++ {
		if A[i].Nama == X {
			idx = i
		}
	}
	return idx
}

func CariSequentialSearch(A manajemenData.TabMakanan, n int) {
	var x string
	fmt.Print("Masukkan nama bahan yang dicari: ")
	fmt.Scan(&x)
	var idx = SequentialSearch(A, n, x)
	if idx != -1 {
		fmt.Printf("Data ditemukan di index ke-%d:\n", idx)
		fmt.Printf("%s - %d buah - Kadaluarsa: %02d-%02d-%04d - Dipakai: %v\n",
			A[idx].Nama, A[idx].JumlahStok, A[idx].TanggalKadaluwarsa, A[idx].BulanKadaluwarsa, A[idx].TahunKadaluwarsa, A[idx].SudahDipakai)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
