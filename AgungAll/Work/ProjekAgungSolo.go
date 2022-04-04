package main

import (
	"fmt"
)

const NMAX int = 1001

type Pelanggan struct {
	nama, gender   string
	usia, transfer int
	saldo          float64
}

type tabPelanggan [NMAX]Pelanggan

type DataPelanggan struct {
	tabel      tabPelanggan
	nPelanggan int
}

func isiDataDiri(data *DataPelanggan) {
	var nama, gender string
	var usia int
	var saldo float64
	data.nPelanggan = 0
	fmt.Scan(&nama, &gender, &usia, &saldo)
	for nama != "NONE" && gender != "NONE" && usia >= 0 && saldo >= 0 {
		data.tabel[data.nPelanggan].nama = nama
		data.tabel[data.nPelanggan].gender = gender
		data.tabel[data.nPelanggan].usia = usia
		data.tabel[data.nPelanggan].saldo = saldo
		fmt.Scan(&nama, &gender, &usia, &saldo)
		data.nPelanggan++
	}
}
func tampilkan(data DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		fmt.Println(data.tabel[i])
	}
}

func totalSaldoPelanggan(data DataPelanggan) float64 {
	var total float64
	for i := 0; i < data.nPelanggan; i++ {
		total += data.tabel[i].saldo
	}
	return total
}

func rata2(data DataPelanggan) float64 {
	return totalSaldoPelanggan(data) / float64(data.nPelanggan)
}

func Minimum(data DataPelanggan) int {
	//asumsi tidak manusia berumur 1000
	var usiaMinimum int = 1000
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].usia < usiaMinimum {
			usiaMinimum = data.tabel[i].usia
		}
	}
	return -1
}
func Maximum(data DataPelanggan) int {
	//asumsi tidak manusia berumur 1000
	var usiaMax int = 1000
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].usia < usiaMax {
			usiaMax = data.tabel[i].usia
		}
	}
	return -1
}

func main() {
	var data DataPelanggan
	var min, max int
	isiDataDiri(&data)
	fmt.Println("")
	tampilkan(data)

	totalSaldo := totalSaldoPelanggan(data)
	fmt.Printf("Saldo Total : %.2f\n", totalSaldo)
	rataSaldo := rata2(data)
	fmt.Printf("Rata-rata Saldo : %.2f\n", rataSaldo)
	min = Minimum(data)
	max = Maximum(data)
	fmt.Println("Usia Minimum dari data :", min)
	fmt.Println("Usia Maximum dari data :", max)

}
