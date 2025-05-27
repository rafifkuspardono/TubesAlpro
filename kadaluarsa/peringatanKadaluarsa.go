package kadaluarsa

import (
	"TubesAlpro/manajemenData"
	"fmt"
)

func PeringatanKadaluarsa(A manajemenData.TabMakanan, n int, hari, bulan, tahun int) {
	fmt.Println("\nBahan makanan yang akan kadaluarsa dalam 7 hari:")
	tidakAda := true

	for i := 0; i < n; i++ {
		if !A[i].SudahDipakai {
			ok, sisa := CekKadaluarsa(hari, bulan, tahun, A[i].TanggalKadaluwarsa, A[i].BulanKadaluwarsa, A[i].TahunKadaluwarsa)
			if ok {
				tidakAda = false
				fmt.Printf("%s - Kadaluarsa: %02d-%02d-%04d (sisa %d hari)\n",
					A[i].Nama, A[i].TanggalKadaluwarsa, A[i].BulanKadaluwarsa, A[i].TahunKadaluwarsa, sisa)
			}
		}
	}

	if tidakAda {
		fmt.Println("Tidak ada bahan yang mendekati tanggal kadaluarsa.")
	}
}
