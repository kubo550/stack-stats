package utils

import (
	"fmt"
	"strconv"
)

func formatNumber(number int) string {
	if number > 10000 {
		//		return fmt.Sprintf("%.1fk", float64(number)/1000)
		return fmt.Sprintf("%dk", number/1000)
	}
	if number > 1000 {
		return fmt.Sprintf("%d", number)[0:1] + "," + fmt.Sprintf("%d", number)[1:]
	}
	return fmt.Sprintf("%d", number)
}

func str(n int) string {
	return strconv.Itoa(n)
}
