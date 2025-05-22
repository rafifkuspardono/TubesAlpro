package main

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func main() {
	var pilihan, nData int
	var makanan manajemenData.TabMakanan

	for {
		menu()
		fmt.Println("Masukkan opsi: \n")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			manajemenData.PrintData(makanan, nData)
		case 2:
			manajemenData.TambahData(&makanan, &nData)
		case 3:
			manajemenData.HapusData(&makanan, &nData)
		}

	}
}

func menu() {
	fmt.Println("\n\t====== APLIKASI MANAJEMEN STOK BAHAN MAKANAN ======")
	fmt.Println("1. Daftar Bahan Makanan")
	fmt.Println("2. Tambah Data Makanan")
	fmt.Println("3. Hapus Data Makanan")
	fmt.Println("0. KELUAR\n")
}
