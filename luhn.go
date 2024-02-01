package main

import (
	"strconv"
)

func validateCard(c int) bool {
	cs := strconv.Itoa(c)
	cLen := len(cs)
	t := 0
	check, _ := ByteToInt(cs[cLen-1])

	seqN := 0

	for i := cLen - 2; i >= 0; i-- {
		d, _ := ByteToInt(cs[i])
		if seqN%2 == 0 {
			d = d * 2
			sd := strconv.Itoa(d)
			d = AddStrDig(sd)
			t += d
		} else {
			t += d
		}
		seqN++
	}
	return check == (10-(t%10))%10
}

func ByteToInt(r byte) (int, error) {
	return strconv.Atoi(string(r))
}

func AddStrDig(sd string) int {
	t := 0
	for i := range sd {
		d, _ := ByteToInt(sd[i])
		t += d
	}
	return t
}
