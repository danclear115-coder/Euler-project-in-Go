package main

import "fmt"

func task2(n int) (sum int) {

	a, b := 1, 1
	c := a + b

	for a < n {
		a = b
		b = c
		c = a + b
		if a % 2 == 0 {
			sum += a
		}
	}
	
	return

}

func main() {
	fmt.Println(task2(4000000))
}
