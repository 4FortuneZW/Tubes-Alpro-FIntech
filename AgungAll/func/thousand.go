package main

import (
	"fmt"
	"strings"
)

func Separate(N interface{}) string {
	switch N.(type) {
	case int:
	case int64:
	case float64:
	case float32:
	default:
		return "Not a valid Number! Doesn't Know!"
	}
	n := fmt.Sprintf("%v", N)
	n = strings.ReplaceAll(n, ",", "")
	dec := ""
	if strings.Index(n, ".") != -1 {
		dec = n[strings.Index(n, ".")+1 : int(len(n))]
		n = n[0:strings.Index(n, ".")]
	}
	for i := 0; i <= len(n); i += 4 {
		a := n[0 : len(n)-i]
		b := n[len(n)-i : int(len(n))]
		n = a + "," + b
	}
	if n[len(n)-1:int(len(n))] == "," {
		n = n[0 : len(n)-1]
	}
	if dec != "" {
		n = n + "." + dec
	}

	return n
}
