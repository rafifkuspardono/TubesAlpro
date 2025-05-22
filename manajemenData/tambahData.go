package manajemenData

import "fmt"

func TambahData(A *TabMakanan, n *int) {
	var i, batas int
	fmt.Println("\nJumlah makanan yang ingin ditambahkan: ")
	fmt.Scan(&batas)
	for i = 0; i < batas; i++ {
		fmt.Println("Masukkan nama makanan: ")
		fmt.Scan(&A[*n].nama)
		fmt.Println("Masukkan jumlah Stok: ")
		fmt.Scan(&A[*n].jumlahStok)
		fmt.Println("Masukkan tanggal kadaluwarsa: ")
		fmt.Scan(&A[*n].tanggalKadaluwarsa)
		fmt.Println("Masukkan bulan kadaluwarsa: ")
		fmt.Scan(&A[*n].bulanKadaluwarsa)
		fmt.Println("Masukkan tahun kadaluwarsa: ")
		fmt.Scan(&A[*n].tahunKadaluwarsa)
		fmt.Println("Status (sudah/belum digunakan)")
		fmt.Scan(&A[*n].sudahDipakai)
		fmt.Println("Data makanan berhasil ditambahkan")
		*n = *n + 1
	}
}
