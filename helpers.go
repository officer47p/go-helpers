package gohelpers

import (
	"fmt"
	"strconv"
)

func FloatToString(n float64) string {
	return fmt.Sprintf("%g", n)
}

func StringToFloat(n string) (float64, error) {
	return strconv.ParseFloat(n, 64)
}
