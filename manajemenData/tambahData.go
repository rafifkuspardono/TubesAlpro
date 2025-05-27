package manajemenData

import "fmt"

func TambahData(A *TabMakanan, n *int) {
	var i, batas int
	fmt.Println("\nJumlah makanan yang ingin ditambahkan: ")
	fmt.Scan(&batas)
	for i = 0; i < batas; i++ {
		fmt.Println("Masukkan nama makanan: ")
		fmt.Scan(&A[*n].Nama)
		fmt.Println("Masukkan jumlah Stok: ")
		fmt.Scan(&A[*n].JumlahStok)
		fmt.Println("Masukkan tanggal kadaluwarsa: ")
		fmt.Scan(&A[*n].TanggalKadaluwarsa)
		fmt.Println("Masukkan bulan kadaluwarsa: ")
		fmt.Scan(&A[*n].BulanKadaluwarsa)
		fmt.Println("Masukkan tahun kadaluwarsa: ")
		fmt.Scan(&A[*n].TahunKadaluwarsa)
		fmt.Println("Status (sudah/belum digunakan)")
		fmt.Scan(&A[*n].SudahDipakai)
		fmt.Println("Data makanan berhasil ditambahkan")
		*n = *n + 1
	}
}
