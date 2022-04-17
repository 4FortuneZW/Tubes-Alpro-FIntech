package main

import (
	"fmt"
)

type Pelanggan struct {
	Nama, JenisKelamin string
	usia               int
	saldo              float64
}

const NMAX = 1001

type tabPelanggan [NMAX]Pelanggan

type DataPelanggan struct {
	tabel      tabPelanggan
	nPelanggan int
}

func DataDiri(data *DataPelanggan, n *int) {
	var nama string
	var usia int
	var saldo float64
	var JenisKelamin string

	for i := 0; i < *n; i++ {
		fmt.Print("Nama : ")
		fmt.Scanln(&nama, &nama, &nama)
		fmt.Print("Usia : ")
		fmt.Scanln(&usia)in)
		fmt.Print("Jenis Kelamin : ")
		fmt.Scanln(&JenisKelam
		fmt.Print("Saldo : ")
		fmt.Scanln(&saldo)

		data.tabel[i].Nama = nama 
		data.tabel[i].usia = usia
		data.tabel[i].JenisKelamin = JenisKelamin
		data.tabel[i].saldo = saldo

	}
}

func findMinAndMax(data DataPelanggan) (min int, max int) {

	min = data.tabel[0].usia
	max = data.tabel[0].usia
	for _, value := range data.tabel { //masuk ke array bukan dari index
		if value.usia < min {
			min = value.usia
		}
		if value.usia > max {
			max = value.usia
		}
	}
	return min, max
}

func main() {
	var data DataPelanggan
	var n int

	fmt.Print("Jumlah orang : ")
	fmt.Scanln(&n)
	DataDiri(&data, &n)
	fmt.Print(data)
	min, max := findMinAndMax(data)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}
