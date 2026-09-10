package main

import "fmt"

func task3() {
	
	number := 600851475143
	divisor := 2

	for divisor*divisor <= number {
		if number%divisor == 0 {
			number /= divisor
		} else {
			divisor++
		}
	}

	fmt.Println(number)

}

func main() {
	fmt.Println(task3)
}
