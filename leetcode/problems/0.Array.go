package problems

import "sort"

//SameArray check two array is same
func SameArray(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func SameStringArray(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	sort.Strings(arr1)
	sort.Strings(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func SameFloat64Array(arr1, arr2 []float64) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
