package main

import (
	"fmt"
)

const NMAX int = 1001

type Pelanggan struct {
	nama, gender string
	usia         int
	saldo        float64
	membership   string
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
		j := i + 1
		for j > 0 && data.tabel[j-1].membership < data.tabel[j].membership {
			temp := data.tabel[j]
			data.tabel[j] = data.tabel[j-1]
			data.tabel[j-1] = temp
			j--
		}
	}
	for i := 0; i < data.nPelanggan; i++ {
		j := i + 1
		for j > 0 && data.tabel[j-1].membership == data.tabel[j].membership && data.tabel[j-1].nama > data.tabel[j].nama {
			temp := data.tabel[j]
			data.tabel[j] = data.tabel[j-1]
			data.tabel[j-1] = temp
			j--
		}
	}
}

func totalSaldoPelanggan(data DataPelanggan) float64 {
	var total float64 = 0.0
	for i := 0; i < data.nPelanggan; i++ {
		total += data.tabel[i].saldo
	}
	return total
}

func rata2Usia(data DataPelanggan) int {
	var total int = 0
	for i := 0; i < data.nPelanggan; i++ {
		total += data.tabel[i].usia
	}
	rata2 := total / data.nPelanggan
	return rata2
}

func rata2SaldoPermembership(data DataPelanggan, member string) float64 {
	//butuh input
	var total float64
	var count int
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].membership == "Silver" && member == "Silver" {
			total += data.tabel[i].saldo
			count++
		} else if data.tabel[i].membership == "Gold" && member == "Gold" {
			total += data.tabel[i].saldo
			count++
		} else if data.tabel[i].membership == "Platinum" && member == "Platinum" {
			total += data.tabel[i].saldo
			count++
		}
	}
	if member == "Silver" || member == "Gold" || member == "Platinum" {
		return total / float64(count)
	}
	return -1
}
func rata2UsiaPermembership(data DataPelanggan, member string) int {
	var total int
	var count int
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].membership == "Silver" && member == "Silver" {
			total += data.tabel[i].usia
			count++
		} else if data.tabel[i].membership == "Gold" && member == "Gold" {
			total += data.tabel[i].usia
			count++
		} else if data.tabel[i].membership == "Platinum" && member == "Platinum" {
			total += data.tabel[i].usia
			count++
		}
	}
	if member == "Silver" || member == "Gold" || member == "Platinum" {
		return total / (count)
	}
	return -1
}

func Minimum(data DataPelanggan) int {
	//asumsi tidak manusia berumur 1000
	var usiaMinimum int = 1000
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].usia < usiaMinimum {
			usiaMinimum = data.tabel[i].usia
		}
	}
	if usiaMinimum >= 0 {
		return usiaMinimum
	}
	return -1
}
func Maximum(data DataPelanggan) int {
	//asumsi tidak manusia berumur -1
	var usiaMax int = -1
	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].usia > usiaMax {
			usiaMax = data.tabel[i].usia
		}
	}
	return usiaMax
}

func tampilkan(data DataPelanggan) {
	for i := 0; i < data.nPelanggan; i++ {
		fmt.Println(data.tabel[i])
	}
}

func counter(data DataPelanggan) (int, int, int) {
	var Silver = 0
	var Gold = 0
	var Platinum = 0

	for i := 0; i < data.nPelanggan; i++ {
		if data.tabel[i].membership == "Silver" {
			Silver += 1
		} else if data.tabel[i].membership == "Gold" {
			Gold += 1
		} else if data.tabel[i].membership == "Platinum" {
			Platinum += 1
		}
	}
	return Silver, Gold, Platinum
}
func main() {
	var data DataPelanggan
	var min, max int
	var Silver, Gold, Platinum int
	isiDataDiri(&data)
	fmt.Println("")
	kelolaMember(&data)
	pengurutanMembership(&data)
	tampilkan(data)
	Silver, Gold, Platinum = counter(data)
	fmt.Println("Jumlah Membership Silver :", Silver)
	fmt.Println("Jumlah Membership Gold :", Gold)
	fmt.Println("Jumlah Membership Platinum :", Platinum)
	totalSaldo := totalSaldoPelanggan(data)
	fmt.Printf("Saldo Total : %.2f\n", totalSaldo)
	min = Minimum(data)
	max = Maximum(data)
	fmt.Println("Usia Minimum dari data :", min)
	fmt.Println("Usia Maximum dari data :", max)
	rata2Usia := rata2Usia(data)
	fmt.Println("Rata-rata usia dari data adalah ", rata2Usia)
	rata2SaldoSilver := rata2SaldoPermembership(data, "Silver")
	rata2SaldoGold := rata2SaldoPermembership(data, "Gold")
	rata2SaldoPlatinum := rata2SaldoPermembership(data, "Platinum")
	rata2UsiaSilver := rata2UsiaPermembership(data, "Silver")
	rata2UsiaGold := rata2UsiaPermembership(data, "Gold")
	rata2UsiaPlatinum := rata2UsiaPermembership(data, "Platinum")
	fmt.Printf("Rata-rata Saldo Silver: %.2f\n", rata2SaldoSilver)
	fmt.Printf("Rata-rata Saldo Gold: %.2f\n", rata2SaldoGold)
	fmt.Printf("Rata-rata Saldo Platinum: %.2f\n", rata2SaldoPlatinum)
	fmt.Printf("Rata-rata Usia Silver: %d\n", rata2UsiaSilver)
	fmt.Printf("Rata-rata Usia Gold: %d\n", rata2UsiaGold)
	fmt.Printf("Rata-rata Usia Platinum: %d\n", rata2UsiaPlatinum)
}
