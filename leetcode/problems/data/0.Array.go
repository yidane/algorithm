package data

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