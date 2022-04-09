package main

import "fmt"

type mahasiswa struct {
	nama        string
	tinggiBadan int
}

const NMAX = 100

type dataMhs [NMAX]mahasiswa

func bacaData(t *dataMhs, n *int) {
	var tinggiBadan int
	var nama string

	for i := 0; i < *n; i++ {
		fmt.Scan(&nama)
		fmt.Scan(&tinggiBadan)

		t[i].nama = nama
		t[i].tinggiBadan = tinggiBadan

	}

}

func main() {
	var t dataMhs
	var n int

	fmt.Scanln(&n)
	bacaData(&t, &n)
	fmt.Print(t)

}
