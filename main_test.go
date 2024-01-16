package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	scenario := []struct {
		input    int
		expected []string
	}{
		{
			input:    1,
			expected: []string{"1"},
		},
		{
			input:    2,
			expected: []string{"1", "2"},
		},
		{
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			input:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			input: 15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for scindex, scenario := range scenario {
		res := FizzBuzz(scenario.input)
		if len(res) != len(scenario.expected) {
			t.Error("FizzBuzz wrong lengths")
		}

		for i := range res {
			if res[i] != scenario.expected[i] {
				t.Errorf("FizzBuzz scen %d range %d expected %s got %s",
					scindex, i, scenario.expected[i], res[i],
				)
			}
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	scenarii := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for scindex, scenario := range scenarii {
		res := ContainsDuplicate(scenario.input)
		if res != scenario.expected {
			t.Errorf("ContainsDuplicate scen %d expected %t got %t",
				scindex, scenario.expected, res,
			)
		}
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

func TestReverseString(t *testing.T) {
	scenarii := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			input:    []byte("Hannah"),
			expected: []byte("hannaH"),
		},
	}

	for scindex, scenario := range scenarii {
		res := ReverseString(scenario.input)
		if len(res) != len(scenario.expected) {
			t.Error("ReverseString wrong lengths")
		}

		for i := range res {
			if res[i] != scenario.expected[i] {
				t.Errorf("ReverseString scen %d range %d expected %d got %d",
					scindex, i, scenario.expected[i], res[i],
				)
			}
		}
	}
}
