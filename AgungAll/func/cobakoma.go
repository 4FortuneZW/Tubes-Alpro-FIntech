package main

import "fmt"

func main() {
	n := 10000.54
	bInt := int(n)
	desimal := (float64(bInt) - n) * -1
	fmt.Printf("%.3f", desimal)

}
