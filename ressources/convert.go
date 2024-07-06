package goreloaded

import (
	"strconv"
	"fmt"
)

func HexConv(hex string) string {
	nb, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(nb)
}

func BinConv(bin string) string {
	nb, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(nb)
}