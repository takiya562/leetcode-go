package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type complexNumber struct {
	real, imag int
}

func complexNumberMultiply(num1 string, num2 string) string {
	n1 := extractComplexNumber(num1)
	n2 := extractComplexNumber(num2)
	real := n1.real*n2.real - n1.imag*n2.imag
	imag := n1.real*n2.imag + n1.imag*n2.real
	return fmt.Sprintf("%d+%di", real, imag)
}

func extractComplexNumber(num string) complexNumber {
	i := strings.Index(num, "+")
	real, _ := strconv.Atoi(num[0:i])
	imag, _ := strconv.Atoi(num[i+1 : len(num)-1])
	return complexNumber{real, imag}
}

func ComplexNumberMultiply(num1 string, num2 string) string {
	return complexNumberMultiply(num1, num2)
}
