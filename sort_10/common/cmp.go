package common_util

type args struct {
	Nums []int
}

var Tests = []struct {
	Name string
	Args args
	Want []int
}{
	// TODO: Add test cases.
	{"case1:", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
	{"case2:", args{[]int{4, 2, 3, 1}}, []int{1, 2, 3, 4}},
	{"case3:", args{[]int{1, 2, 6, 4, 7, 8, 9}}, []int{1, 2, 4, 6, 7, 8, 9}},
}

func SwapInt(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

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
