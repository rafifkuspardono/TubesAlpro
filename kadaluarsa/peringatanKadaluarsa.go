package kadaluarsa

import "fmt"

func peringatanKadaluarsa(A TabMakanan, n int, hari, bulan, tahun int) {
	fmt.Println("\nBahan makanan yang akan kadaluarsa dalam 7 hari:")
	tidakAda := true

	for i := 0; i < n; i++ {
		if !data[i].sudahDipakai {
			ok, sisa := dalam7Hari(hari, bulan, tahun, A[i].tanggalKadaluwarsa, A[i].bulanKadaluwarsa, A[i].tahunKadaluwarsa)
			if ok {
				tidakAda = false
				fmt.Printf("%s - Kadaluarsa: %02d-%02d-%04d (sisa %d hari)\n",
					A[i].nama, A[i].tanggalKadaluwarsa, A[i].bulanKadaluwarsa, A[i].tahunKadaluwarsa, sisa)
			}
		}
	}

	if tidakAda {
		fmt.Println("Tidak ada bahan yang mendekati tanggal kadaluarsa.")
	}
}
