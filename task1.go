package main

import "fmt"

func task1(n int) (sum int) {

	for num := 0; num < n; num++ {
		if num % 3 == 0 || num % 5 == 0 {
			sum += num
		} 
	}

	return

}

func main() {
	fmt.Println(task1(1000))
}
