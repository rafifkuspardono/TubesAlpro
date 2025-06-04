package manajemenData

type BahanMakanan struct {
	Nama                                                   string
	JumlahStok                                             int
	TanggalKadaluwarsa, BulanKadaluwarsa, TahunKadaluwarsa int
	SudahDipakai                                           bool
}

type TabMakanan [100]BahanMakanan
