package main

import "fmt"

type bahanMakanan struct {
	nama                                                   string
	jumlahStok                                             int
	tanggalKadaluwarsa, bulanKadaluwarsa, tahunKadaluwarsa int
	sudahDipakai                                           bool
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
