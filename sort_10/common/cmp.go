package common_util

//Use to compare slice int type
func IntSliceEqual(a, b []int) bool {
	//len cmp
	if len(a) != len(b) {
		return false
	}
	//nil cmp
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
