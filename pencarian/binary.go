package pencarian

import (
	"TubesAlpro/manajemenData"
	"TubesAlpro/pengurutan"
	"fmt"
)

func BinarySearchNama(A manajemenData.TabMakanan, n int, x string) int {
	var low, high, mid int
	low = 0
	high = n - 1

	for low <= high {
		mid = (low + high) / 2
		if A[mid].Nama == x {
			return mid
		} else if x < A[mid].Nama {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func CariBinary(A *manajemenData.TabMakanan, n int) {
	var x string
	fmt.Println("Masukkan nama bahan yang dicari (Binary Search): ")
	fmt.Scan(&x)

	pengurutan.SelectionNama(A, n)

	idx := BinarySearchNama(*A, n, x)
	if idx != -1 {
		fmt.Printf("%s - %d buah - Kadaluarsa: %02d-%02d-%04d - Dipakai: %v\n",
			A[idx].Nama, A[idx].JumlahStok, A[idx].TanggalKadaluwarsa, A[idx].BulanKadaluwarsa, A[idx].TahunKadaluwarsa, A[idx].SudahDipakai)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
