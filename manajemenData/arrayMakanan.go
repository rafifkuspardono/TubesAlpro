package manajemenData

type bahanMakanan struct {
	nama                                                   string
	jumlahStok                                             int
	tanggalKadaluwarsa, bulanKadaluwarsa, tahunKadaluwarsa int
	sudahDipakai                                           bool
}

type TabMakanan [100]bahanMakanan
