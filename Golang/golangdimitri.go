func kelolaMembership(data *DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].saldo <= 25000000.0 {
			data.tabel[i].membership = "Silver"
		} else if (data.tabel[i].saldo > 25000000.0) && (data.tabel[i].saldo <= 100000000.0) {
			data.tabel[i].membership = "Gold"
		} else if data.tabel[i].saldo > 100000000.0 {
			data.tabel[i].membership = "Platinum"
		}
	}
}

func pengurutanMembership(data *DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		j := i + 1
		for (data.tabel[j-1].membership < data.tabel[j].membership) && (j > 0) {
			temp := data.tabel[j]
			data.tabel[j] = data.tabel[j-1]
			data.tabel[j-1] = temp
			j--
		}
	}
	for i := 0; i < data.nPelanggan; i++ {
		j := i + 1
		for (data.tabel[j-1].membership == data.tabel[j].membership) && (data.tabel[j-1].nama > data.tabel[j].nama) && (j > 0) {
			temp := data.tabel[j]
			data.tabel[j] = data.tabel[j-1]
			data.tabel[j-1] = temp
			j--
		}
	}
}

