package main

import (
	"strconv"
)

func main() {

}

func FizzBuzz(n int) []string {
	answers := make([]string, n+1)
	for i := 1; i < n+1; i++ {
		if i%15 == 0 {
			answers[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			answers[i-1] = "Fizz"
		} else if i%5 == 0 {
			answers[i-1] = "Buzz"
		} else {
			answers[i-1] = strconv.Itoa(i)
		}
	}

	return answers
}

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func MovingZeroes(nums []int) []int {
	zeroesCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroesCount++
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	nums = append(nums, make([]int, zeroesCount)...)

	return nums
}
