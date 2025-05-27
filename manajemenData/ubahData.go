package manajemenData

import "fmt"

func UbahData(A *TabMakanan, n int) {
	var i int
	var nama string
	fmt.Print("Masukkan nama bahan yang akan diubah: ")
	fmt.Scan(&nama)
	for i = 0; i < n; i++ {
		if A[i].Nama == nama {
			fmt.Println("Nama baru: ")
			fmt.Scan(&A[i].Nama)
			fmt.Println("Jumlah stok baru: ")
			fmt.Scan(&A[i].JumlahStok)
			fmt.Println("Tanggal kadaluwarsa baru: ")
			fmt.Scan(&A[i].TanggalKadaluwarsa)
			fmt.Println("Bulan kadaluwarsa baru: ")
			fmt.Scan(&A[i].BulanKadaluwarsa)
			fmt.Println("Tahun kadaluwarsa baru: ")
			fmt.Scan(&A[i].TahunKadaluwarsa)
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
