package manajemenData

type bahanMakanan struct {
	Nama                                                   string
	JumlahStok                                             int
	TanggalKadaluwarsa, BulanKadaluwarsa, TahunKadaluwarsa int
	SudahDipakai                                           bool
}

type TabMakanan [100]bahanMakanan
