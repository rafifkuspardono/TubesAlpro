package manajemenData

import "fmt"

func PrintData(A TabMakanan, n int) {
	if n == 0 {
		fmt.Println("\nData masih kosong.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %d buah - Kadaluarsa: %02d-%02d-%04d - Dipakai: %v\n",
			i+1, A[i].Nama, A[i].JumlahStok, A[i].TanggalKadaluwarsa, A[i].BulanKadaluwarsa, A[i].TahunKadaluwarsa, A[i].SudahDipakai)
	}
}
