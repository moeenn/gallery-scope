package record

import (
	"testing"
)

func TestRecordNormalize(t *testing.T) {
	testCases := []struct {
		Input  Record
		Output Record
	}{
		{
			Input:  Record{Base: "http://www.google.com/", Start: 1, End: 15},
			Output: Record{Base: "http://www.google.com/", Start: 1, End: 15},
		},
		{
			Input:  Record{Base: "http://www.google.com/", End: 15},
			Output: Record{Base: "http://www.google.com/", Start: 1, End: 15, Zero: 0},
		},
		{
			Input:  Record{Base: "http://www.google.com/", End: 15, Zero: 3},
			Output: Record{Base: "http://www.google.com/", Start: 1, End: 15, Zero: 3},
		},
	}

	isSame := func(x, y Record) bool {
		return x.Base == y.Base &&
			x.Start == y.Start &&
			x.End == y.End &&
			x.Zero == y.Zero
	}

	for _, testCase := range testCases {
		testCase.Input.Normalize()

		if !isSame(testCase.Input, testCase.Output) {
			t.Errorf("Failed:: Expected: %v Got: %v", testCase.Output, testCase.Input)
		}
	}
}

func TestRecordProcess(t *testing.T) {
	testCases := []struct {
		Input    Record
		Expected []string
	}{
		{
			Input: Record{
				Base:  "https://www.google.com/12345/img_{n}",
				Start: 1,
				End:   4,
			},
			Expected: []string{
				"https://www.google.com/12345/img_1",
				"https://www.google.com/12345/img_2",
				"https://www.google.com/12345/img_3",
				"https://www.google.com/12345/img_4",
			},
		},
		{
			Input: Record{
				Base: "https://www.google.com/12345/img_{n}",
				End:  3,
				Zero: 3,
			},
			Expected: []string{
				"https://www.google.com/12345/img_001",
				"https://www.google.com/12345/img_002",
				"https://www.google.com/12345/img_003",
			},
		},
	}

	for _, testCase := range testCases {
		got, err := testCase.Input.Process()
		if err != nil {
			t.Errorf("[failed] Failed to process record: %v", err)
		}

		if len(got) != len(testCase.Expected) {
			t.Errorf("[failed] Actual URL array length (i.e. %d) different from expected length (i.e. %d)", len(got), len(testCase.Expected))
		}

		for i := 0; i < len(got); i++ {
			if got[i] != testCase.Expected[i] {
				t.Errorf("[failed] Expected: %s Got: %s\n", testCase.Expected[i], got[i])
			}
		}
	}
}
