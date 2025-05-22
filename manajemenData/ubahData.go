package manajemenData

import "fmt"

func UbahData(A *TabMakanan, n int) {
	var i int
	var nama string
	fmt.Print("Masukkan nama bahan yang akan diubah: ")
	fmt.Scan(&nama)
	for i = 0; i < n; i++ {
		if A[i].nama == nama {
			fmt.Println("Nama baru: ")
			fmt.Scan(&A[i].nama)
			fmt.Println("Jumlah stok baru: ")
			fmt.Scan(&A[i].jumlahStok)
			fmt.Println("Tanggal kadaluwarsa baru: ")
			fmt.Scan(&A[i].tanggalKadaluwarsa)
			fmt.Println("Bulan kadaluwarsa baru: ")
			fmt.Scan(&A[i].bulanKadaluwarsa)
			fmt.Println("Tahun kadaluwarsa baru: ")
			fmt.Scan(&A[i].tahunKadaluwarsa)
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
