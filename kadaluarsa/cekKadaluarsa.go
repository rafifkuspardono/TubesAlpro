package kadaluarsa

func CekKadaluarsa(h1, b1, t1, h2, b2, t2 int) (bool, int) {
	selisih := HitungJumlahHari(h2, b2, t2) - HitungJumlahHari(h1, b1, t1)
	return selisih >= 0 && selisih <= 7, selisih
}
