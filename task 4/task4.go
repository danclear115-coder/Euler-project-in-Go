package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {

	strNum, reverseStrNum := strconv.Itoa(n), ""

	for i := len(strNum) - 1; i >= 0; i-- {
		reverseStrNum += string(strNum[i])
	}

	reverseInt, err := strconv.Atoi(reverseStrNum)

	if err != nil {
		fmt.Println(err)
	}

	return n == reverseInt

}

func task4() (maxNum int) {
	
	for a := 100; a <= 999; a++ {
		for b := 100; b <= 990; b++ {
			multNum := a * b
			if isPalindrome(multNum) && maxNum < multNum {
				maxNum = multNum
			}
		}
	}

	return maxNum

}

func main() {
	fmt.Println(task4())
}
