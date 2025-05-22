package main

import "fmt"

type bahanMakanan struct {
	nama, kadaluwarsa string
	stok int 
}

type arrBahan [100]bahanMakanan

func main() {
	var makanan bahanMakanan
	
	menu()
}

func menu() {
	fmt.Println("\n		====== APLIKASI MANAJEMEN STOK BAHAN MAKANAN ======\n")
	fmt.Println("1. Tambah Bahan Makanan")
	fmt.Println("2. Ubah Data Bahan Makanan")
	fmt.Println("3. Hapus Data Bahan Makanan")
	fmt.Println("4. Mencari Bahan Makanan (sequential)")
	fmt.Println("5. Mencari Bahan Makanan (Binary)")
	fmt.Println("6. Mengurutkan Berdasarkan Tanggal Kadaluwarsa (Selection)")
	fmt.Println("7. Mengurutkan Berdasarkan Jumlah Stok (Insertion)")
	fmt.Println("8. Apakah Bahan Makanan Mendekati Expired?")
	fmt.Println("9. Total Bahan Makanan yang Tersedia")
	fmt.Println("10. Total Bahan Makanan yang telah Digunakan\n")
	fmt.Println("0. KELUAR\n")

}

func tambahData(list *arrBahan, n *int) {
	
}
