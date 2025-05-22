package manajemenData

import "fmt"

func PrintData(A TabMakanan, n int) {
	if n == 0 {
		fmt.Println("\nData masih kosong.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %d buah - Kadaluarsa: %02d-%02d-%04d - Dipakai: %v\n",
			i+1, A[i].nama, A[i].jumlahStok, A[i].tanggalKadaluwarsa, A[i].bulanKadaluwarsa, A[i].tahunKadaluwarsa, A[i].sudahDipakai)
	}
}
