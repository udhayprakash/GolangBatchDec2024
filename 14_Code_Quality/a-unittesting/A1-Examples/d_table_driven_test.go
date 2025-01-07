package main

import "testing"

type TestCase struct {
	inputSlice []int
	expected   int
}

func TestMax5(t *testing.T) {
	cases := []TestCase{
		TestCase{
			inputSlice: []int{1, 2, 3, 4, 5},
			expected:   5,
		},
		TestCase{
			inputSlice: []int{1, 2, 3, 4, 5},
			expected:   6,
		},
		TestCase{
			inputSlice: []int{},
			expected:   0,
		},
	}
	for _, c := range cases {
		actual := Max2(c.inputSlice)
		expected := c.expected
		if actual != expected {
			// t.Logf("Expected %d, got %d", expected, actual)
			// t.FailNow()

			// t.Fatalf() function, which is a combination of t.FailNow() + t.Logf()
			t.Fatalf("Expected %d, got %d", expected, actual)
		}
	}
}

func Max2(numbers []int) int {
	var max int
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

/*
NOTE: Max function can be used, when running all test cases togther only


 $ go test -v d_table_driven_test.go
=== RUN   TestMax5
    d_table_driven_test.go:29: Expected 6, got 5
--- PASS: TestMax5 (0.00s)
PASS
ok      command-line-arguments  0.002s
@udhayprakash ➜ .../GolangBatchDec2024/14_Code_Quality/a-unittesting/A1-Examples (Day10) $ go test -v d_table_driven_test.go
=== RUN   TestMax5
    d_table_driven_test.go:29: Expected 6, got 5
--- FAIL: TestMax5 (0.00s)
FAIL
FAIL    command-line-arguments  0.001s
FAIL
*/