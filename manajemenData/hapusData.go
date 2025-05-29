package manajemenData

import "fmt"

func HapusData(A *TabMakanan, n *int) {
	var i, j int
	var nama string
	fmt.Print("Masukkan nama bahan yang akan dihapus: ")
	fmt.Scan(&nama)
	for i = 0; i < *n; i++ {
		if A[i].Nama == nama {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
