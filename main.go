package main

import (
	"strconv"
)

func main() {

}

func FizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res

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

func ReverseString(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	return s
}
