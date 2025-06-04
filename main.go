package main

import (
	"TubesAlpro/kadaluarsa"
	"TubesAlpro/manajemenData"
	"TubesAlpro/pencarian"
	"TubesAlpro/pengurutan"
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
		case 4:
			manajemenData.UbahData(&makanan, nData)
		case 5:
			var hariSekarang, bulanSekarang, tahunSekarang int
			fmt.Println("Masukkan tanggal hari ini: ")
			fmt.Scan(&hariSekarang, &bulanSekarang, &tahunSekarang)
			fmt.Println("Masukkan bulan hari ini: ")
			fmt.Scan(&bulanSekarang)
			fmt.Println("Masukkan tahun hari ini: ")
			fmt.Scan(&tahunSekarang)
			kadaluarsa.PeringatanKadaluarsa(makanan, nData, hariSekarang, bulanSekarang, tahunSekarang)
		case 6:
			pencarian.CariSequentialSearch(makanan, nData)
		case 7:
			pengurutan.SelectionSortKadaluarsa(&makanan, nData)
			manajemenData.PrintData(makanan, nData)
		case 8:
			pengurutan.InsertionSortStok(&makanan, nData)
			manajemenData.PrintData(makanan, nData)
		}

	}
}

func menu() {
	fmt.Println("\n\t====== APLIKASI MANAJEMEN STOK BAHAN MAKANAN ======")
	fmt.Println("1. Daftar Bahan Makanan")
	fmt.Println("2. Tambah Data Makanan")
	fmt.Println("3. Hapus Data Makanan")
	fmt.Println("4. Ubah Data Makanan")
	fmt.Println("5. Peringatan Kadaluarsa")
	fmt.Println("6. Cari Data (Sequential)")
	fmt.Println("7. Cari Data (Binary)")
	fmt.Println("8. Pengurutan Data (Selection)")
	fmt.Println("9. Pengurutan Data (Insertion)")
	fmt.Println("0. KELUAR\n")
}
