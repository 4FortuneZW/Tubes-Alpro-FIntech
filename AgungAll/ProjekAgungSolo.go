package main

import "fmt"

const NMAX int = 1001

type Pelanggan struct {
	nama, gender          string
	usia, saldo, transfer int
}

type tabPelanggan [NMAX]Pelanggan

type DataPelanggan struct {
	tabel      tabPelanggan
	nPelanggan int
}

func isiDataDiri(data *DataPelanggan) {
	var nama, gender string
	var usia, saldo int
	data.nPelanggan = 0
	fmt.Scan(&nama, &gender, &usia, &saldo)
	for nama != "NONE" && gender != "NONE" && usia >= 0 && saldo >= 0 {
		data.tabel[data.nPelanggan].nama = nama
		data.tabel[data.nPelanggan].gender = gender
		data.tabel[data.nPelanggan].usia = usia
		data.tabel[data.nPelanggan].saldo = saldo
		fmt.Scan(&nama, &gender, &usia, &saldo)
	}
}
func tampilkan(data DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		fmt.Println(data.tabel[i])
	}
}

func main() {

	fmt.Println("YUK BISA YUK")
}
