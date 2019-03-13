package problems

import "testing"

func Test_maxProduct(t *testing.T) {
	type testCase struct {
		arr    []string
		result int
	}

	testCases := []testCase{
		{arr: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}, result: 16},
		{arr: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, result: 4},
		{arr: []string{"a", "aa", "aaa", "aaaa"}, result: 0},
	}

	for _, tc := range testCases {
		result := maxProduct(tc.arr)
		if result != tc.result {
			t.Fatalf("%s should return %v", tc.arr, tc.result)
		}
	}
}
