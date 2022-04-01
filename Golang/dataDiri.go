package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Nama, JenisKelamin string
	usia int
	saldo float64
}
const NMAX = 30

type tabUser  [NMAX]User

func DataDiri(DtUser *tabUser, n *int) {
	var nama string
	var usia int
	var saldo float64
	var JenisKelamin string
	
	
	fmt.Scanln(&n)
	i := 0
	for i < n && *DtUser <= NMAX-1 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		nama := scanner.Text()
		fmt.Println("Nama :", nama)
		fmt.Print("Usia : ")
		fmt.Scanln(&usia)
		fmt.Print("Jenis Kelamin : ")
		fmt.Scanln(&JenisKelamin)
		fmt.Print("Saldo : ")
		fmt.Scanln(&saldo)

		tabUser[i].Nama = nama
		tabUser[i].Usia = usia
		tabUser[i].Jenis_Kelamin = JenisKelamin
		tabUser[i].Saldo = saldo
		*DtUser++
		i++
}  



func main() {
	var Datadiri tabUser
	var n,
	


}
