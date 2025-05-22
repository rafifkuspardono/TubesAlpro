package main

import "fmt"

type bahanMakanan struct {
	nama                                                   string
	jumlahStok                                             int
	tanggalKadaluwarsa, bulanKadaluwarsa, tahunKadaluwarsa int
}

type tabMakanan [100]bahanMakanan

func main() {
	var pilihan, nData int
	var makanan tabMakanan

	for {
		menu()
		fmt.Println("Masukkan opsi: \n")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			printData(makanan, nData)
		case 2:
			tambahData(&makanan, &nData)
		}

	}
}

func menu() {
	fmt.Println("\n\t====== APLIKASI MANAJEMEN STOK BAHAN MAKANAN ======")
	fmt.Println("1. Daftar Bahan Makanan")
	fmt.Println("2. Tambah Data Makanan")
	fmt.Println("0. KELUAR\n")
}

func tambahData(A *tabMakanan, n *int) {
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
		fmt.Println("Data makanan berhasil ditambahkan")
		*n = *n + 1
	}
}
