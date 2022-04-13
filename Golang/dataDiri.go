@@ -1,68 +0,0 @@
package main

import (
	"fmt"
)

type User struct {
	Nama, JenisKelamin string
	usia               int
	saldo              float64
}

const NMAX = 4

type tabUser [NMAX]User

func DataDiri(DtUser *tabUser, n *int) {
	var nama string
	var usia int
	var saldo float64
	var JenisKelamin string

	for i := 0; i < *n; i++ {
		fmt.Print("Nama : ")
		fmt.Scanln(&nama, &nama, &nama)
		fmt.Print("Usia : ")
		fmt.Scanln(&usia)
		fmt.Print("Jenis Kelamin : ")
		fmt.Scanln(&JenisKelamin)
		fmt.Print("Saldo : ")
		fmt.Scanln(&saldo)

		DtUser[i].Nama = nama //
		DtUser[i].usia = usia
		DtUser[i].JenisKelamin = JenisKelamin
		DtUser[i].saldo = saldo

	}
}

func findMinAndMax(DtUser tabUser) (min int, max int) {

	min = DtUser[0].usia
	max = DtUser[0].usia
	for _, value := range DtUser { //masuk ke array bukan dari index
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
	var DtUser tabUser
	var n int

	fmt.Print("Jumlah orang : ")
	fmt.Scanln(&n)
	DataDiri(&DtUser, &n)
	fmt.Print(DtUser)
	min, max := findMinAndMax(DtUser)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}