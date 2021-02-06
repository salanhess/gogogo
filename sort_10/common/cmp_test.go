package common_util

import "testing"

func TestIntSliceEqual(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//对于一些复合数据类型，在使用时应该声明其类型信息，不然Go无法自动判别报错missing type in composite literal
		{"len cmp1:", args{[]int{1, 2}, []int{3, 4}}, false},
		{"len cmp1:", args{[]int{1, 23, 4}, []int{1, 23, 4}}, true},
		{"len cmp2:", args{[]int{1, 2}, []int{3}}, false},
		{"len cmp3:", args{[]int{}, []int{}}, true},
		{"len cmp4:", args{[]int{1, 2}, []int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSliceEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntSliceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapInt(t *testing.T) {
	type args struct {
		a *int
		b *int
	}
	a, b := 1, 2
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{&a, &b}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
