package main 


import "fmt"

const NMAX = 2021

type tabInt [NMAX]int

func minimum(T tabInt, N int) int {
	min := T[0]
	k := 1
	for k < N {
		if min > T[k] {
			min = T[k]
		}
		k += 1
	}
	return min
}

func maximum(T tabInt, N int) int {
	max := T[0]
	k := 1
	for k < N {
		if max < T[k] {
			max = T[k]
		}
		k += 1
	}
	return max
}

func main() {
	var T = tabInt{9, 21, 4, 13, 29, 24, 5, 48, 33, 37, 3, 17, 20, 40}
	nilai_minimum := minimum(T, 13)
	nilai_maximum := maximum(T, 13)
	fmt.Println(nilai_minimum)
	fmt.Println(nilai_maximum)
}

/*
function minimum(T: tabInt, N:integer)  integer
{Diterima array T yang berisi N bilangan bulat untuk mengembalikan nilai minimum dari array T}
kamus
    min : integer  {variable untuk nilai ekstrim}
    k : integer
algoritma
    min <- T[0]    {data pertama ada di indeks ke-0}
    k <- 1         {perbandingan dilakukan dari data ke-2 hingga N-1}
    while k < N do
        if min > T[k] then     {cek apabila nilai ekstrim tidak valid}
            min <- T[k]        {update nilai ekstrim dengan yang valid}
        endif
        k <- k + 1
    endwhile
    return min       {setelah semua data dicek, maka nilai ekstrim valid}
endfunction
*/
