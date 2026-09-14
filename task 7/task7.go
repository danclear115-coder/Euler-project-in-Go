package main

import "fmt"

func task7() int {

	mulSum, sumOfMul := 0, 0

	for i := 1; i <= 100; i++ {
		mulSum += i
		sumOfMul += i * i
		fmt.Println(i, i * i, sumOfMul)
	}

	fmt.Println(mulSum, sumOfMul)

	return ((mulSum * mulSum) - sumOfMul)

}

func main() {
	fmt.Println(task7())
}
