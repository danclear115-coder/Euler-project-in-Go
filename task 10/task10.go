package main

import "fmt"

func isSimple(n int) bool {

	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true

}

func task10() int {

	sum := 0

	for i := 2; i < 10; i++ {
		if isSimple(i) {
			sum += i
		}
	}

	return sum

}

func main() {
	fmt.Println(task10())
}
