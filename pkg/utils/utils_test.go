package utils

import (
	"testing"
)

func TestZeroPad(t *testing.T) {
	testCases := []struct {
		Input  int
		Len    int
		Output string
	}{
		{Input: 5, Len: 0, Output: "5"},
		{Input: 5, Len: 3, Output: "005"},
		{Input: 10, Len: 0, Output: "10"},
		{Input: 20, Len: 5, Output: "00020"},
	}

	for _, testCase := range testCases {
		got := ZeroPad(testCase.Input, testCase.Len)
		if got != testCase.Output {
			t.Errorf("[failed] Expected: %s Got: %s", testCase.Output, got)
		}
	}
}

func TestRandomStr(t *testing.T) {
	testCases := []int{5, 10, 15, 50}
	for _, testCase := range testCases {
		got, err := RandomStr(testCase)

		if err != nil {
			t.Errorf("[failed] Generating random string threw error %v", err)
		}

		if len(got) != testCase {
			t.Errorf("[failed] Length of the random string different from desired length")
		}
	}
}

func TestInsertCount(t *testing.T) {
	testCases := []struct {
		URL, Count, Expected string
	}{
		{
			URL:      "https://www.google.com/123123/12313_{n}.jpg",
			Count:    "10",
			Expected: "https://www.google.com/123123/12313_10.jpg",
		},
		{
			URL:      "https://www.google.com/123123-{n}/12313_{n}.jpg",
			Count:    "007",
			Expected: "https://www.google.com/123123-007/12313_007.jpg",
		},
	}

	for _, testCase := range testCases {
		got := InsertCount(testCase.URL, testCase.Count)
		if got != testCase.Expected {
			t.Errorf("[failed] Expected: %s Got: %s\n", testCase.Expected, got)
		}
	}
}
