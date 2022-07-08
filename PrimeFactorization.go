package main

import (
	"fmt"
	"strconv"
)

// Get all prime factors of a given number n
func PrimeFactors(n int) (slc []int) {
	for n%2 == 0 {
		slc = append(slc, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			slc = append(slc, i)
			n = n / i
		}
	}
	if n > 2 {
		slc = append(slc, n)
	}
	return slc
}

func main() {
	slc := []int{}
	str, str2 := "", ""
	var number int = 20

	slc = PrimeFactors(number)
	fmt.Println(len(slc))

	for i := 0; i < len(slc); i++ {
		str += strconv.Itoa(slc[i]) + " x "
	}
	str2 += str[:9] + " = " + strconv.Itoa(number)
	fmt.Println(str2)
}
