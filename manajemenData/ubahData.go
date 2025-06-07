package manajemenData

import "fmt"

func UbahData(T *TabMakanan, n int) {
	var nama string
	var i, pilihan int

	fmt.Print("Masukkan nama bahan yang akan diubah: ")
	fmt.Scan(&nama)

	for i = 0; i < n; i++ {
		if T[i].Nama == nama {
			fmt.Println("Data ditemukan:")
			fmt.Printf("1. Ubah Nama\n2. Ubah Jumlah Stok\n3. Ubah Tanggal Kadaluarsa\n4. Tandai Sudah Dipakai\nPilihan: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Print("Nama baru: ")
				fmt.Scan(&T[i].Nama)
			case 2:
				fmt.Print("Jumlah stok baru: ")
				fmt.Scan(&T[i].JumlahStok)
			case 3:
				fmt.Print("Tanggal kadaluarsa baru (DD MM YYYY): ")
				fmt.Scan(&T[i].TanggalKadaluwarsa, &T[i].BulanKadaluwarsa, &T[i].TahunKadaluwarsa)
			case 4:
				T[i].SudahDipakai = true
				fmt.Println("Bahan ditandai sebagai sudah dipakai.")
			default:
				fmt.Println("Pilihan tidak valid.")
			}
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}
