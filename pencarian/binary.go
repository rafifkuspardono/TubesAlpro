package pencarian

import (
	"TubesAlpro/manajemenData"
	"TubesAlpro/pengurutan"
	"fmt"
)

func BinarySearch(A manajemenData.TabMakanan, n int, X string) int {
	var left, right, mid int
	var found int = -1
	left = 0
	right = n - 1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if X < A[mid].Nama {
			right = mid - 1
		} else if X > A[mid].Nama {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func CariBinarySearch(A *manajemenData.TabMakanan, n int) {
	pengurutan.UrutkanNamaMakanan(A, n)

	var x string
	fmt.Print("Masukkan nama bahan yang dicari: ")
	fmt.Scan(&x)

	var idx int
	idx = BinarySearch(*A, n, x)

	if idx != -1 {
		fmt.Printf("Data ditemukan di index ke-%d:\n", idx)
		fmt.Printf("%s - %d buah - Kadaluarsa: %02d-%02d-%04d - Dipakai: %v\n",
			A[idx].Nama, A[idx].JumlahStok, A[idx].TanggalKadaluwarsa, A[idx].BulanKadaluwarsa, A[idx].TahunKadaluwarsa, A[idx].SudahDipakai)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
