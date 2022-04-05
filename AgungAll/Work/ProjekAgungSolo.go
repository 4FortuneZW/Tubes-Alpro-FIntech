package main

import (
	"fmt"
)

const NMAX int = 1001

type Pelanggan struct {
	nama, gender, membership string
	usia                     int
	saldo                    float64
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

func kelolaMember(data *DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].saldo <= 25000000.0 {
			data.tabel[i].membership = "Silver"
		} else if data.tabel[i].saldo <= 100000000.0 {
			data.tabel[i].membership = "Gold"
		} else if data.tabel[i].saldo > 100000000.0 {
			data.tabel[i].membership = "Platinum"
		}
	}
}

func pengurutanMembership(data *DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {

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

func rata2Usia(data DataPelanggan) int {
	var total int = 0.0
	for i := 0; i < data.nPelanggan; i++ {
		total += data.tabel[i].usia
	}
	rata2 := total / data.nPelanggan
	return rata2
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

	totalSaldo := totalSaldoPelanggan(data)
	fmt.Printf("Saldo Total : %.2f\n", totalSaldo)
	rataSaldo := rata2(data)
	fmt.Printf("Rata-rata Saldo : %.2f\n", rataSaldo)
	min = Minimum(data)
	max = Maximum(data)
	fmt.Println("Usia Minimum dari data :", min)
	fmt.Println("Usia Maximum dari data :", max)
	rata2Usia := rata2Usia(data)
	fmt.Println("Rata-rata usia dari data adalah ", rata2Usia)

}
