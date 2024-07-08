package goreloaded

import (
	"strconv"
)

func HexConv(hex string) string {
	nb, err := strconv.ParseInt(hex, 16, 64)
	if err == nil {
		return strconv.Itoa(int(nb))
	} else {
		return hex
	}
}

func BinConv(bin string) string {
	nb, err := strconv.ParseInt(bin, 2, 64)
	if err == nil {
		return strconv.Itoa(int(nb))
	} else {
		return bin
	}
}
