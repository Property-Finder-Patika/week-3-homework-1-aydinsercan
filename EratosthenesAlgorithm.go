package main

import "fmt"

func main() {
	FindAllPrimeNum(20)
}

func FindAllPrimeNum(number int) {

	//Creating boolean array
	s := make([]bool, number+1)

	//fmt.Println(len(s),s)

	for i := 2; i*i <= number; i++ {
		if s[i] == false {
			for j := i * 2; j <= number; j += i {
				s[j] = true
			}
		}
	}

	for i := 2; i <= number; i++ {
		if s[i] == false {
			fmt.Printf("%d ", i)
		}
	}

}
