package main

import "fmt"

func main() {
	var nums [10]int

	for i:= 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("Number %d is even\n", num)
		}	else {
			fmt.Printf("Number %d is odd\n", num)
		}
	}
}