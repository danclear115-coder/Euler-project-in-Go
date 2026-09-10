package main

import "fmt"

func func5() int {

	num := 1

	for {
		flag := true

		for i := 1; i <= 20; i++ {

			if num%i != 0 {
				num++
				flag = false
				break
			}
		}

		if flag {
			return num
		}
	}

}

func main() {
	fmt.Println(func5())
}
