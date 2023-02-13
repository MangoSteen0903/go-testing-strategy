package main

import "fmt"

func main() {
	_, result := isPrime(-1)
	fmt.Println(result)
}

func isPrime(num int) (bool, string) {
	result := ""
	if num == 0 || num == 1 {
		result = fmt.Sprintf("%d is not prime number. prime number should be bigger than 1.", num)
		return false, result
	}

	if num < 0 {
		result = fmt.Sprintf("%d : should not be negative number. please try again.", num)
		return false, result
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			result = fmt.Sprintf("%d is not a prime number.", num)
			return false, result
		}
	}
	result = fmt.Sprintf("%d is a prime number.", num)
	return true, result
}
