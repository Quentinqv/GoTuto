package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"}
	actual := FizzBuzz(13)

	for i := 0; i < 5; i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected %s, got %s", expected[i], actual[i])
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	actual := ContainsDuplicate([]int{1, 2, 3, 4, 3})

	if true != actual {
		t.Errorf("Expected %t, got %t", true, actual)
	}
}

func TestMovingZeroes(t *testing.T) {
	scenarii := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{0, 0, 0, 1},
			expected: []int{1, 0, 0, 0},
		},
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{0, 1, 0, 33, 12},
			expected: []int{1, 33, 12, 0, 0},
		},
	}

	for scindex, scenario := range scenarii {
		res := MovingZeroes(scenario.input)
		if len(res) != len(scenario.expected) {
			t.Error("Moving Zeroes wrong lengths")
		}

		for i := range res {
			if res[i] != scenario.expected[i] {
				t.Errorf("MovingZeroes scen %d range %d expected %d got %d",
					scindex, i, scenario.expected[i], res[i],
				)
			}
		}
	}
}
