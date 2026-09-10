package main

import "fmt"

func isSimpleCheck(n int) bool {

	for i := 2; i < n; i++ {
		if n % i == 0 {
			return true
		}
	}

	return false

}

func task3(n int) (div int) {

	for i := 2; i < n; i++ {
		if n % i == 0 && !isSimpleCheck(i) && div < i {
			div = i
		}
	}

	return div

}

func main() {
	fmt.Println(task3(13195))
}
