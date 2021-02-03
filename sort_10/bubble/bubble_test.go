package main

import (
	"reflect"
	"testing"
)

func Test_sort_bubble(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1:", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"case2:", args{[]int{4, 2, 3, 1}}, []int{1, 2, 3, 4}},
		{"case3:", args{[]int{11, 12, 3, 14}}, []int{3, 11, 12, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort_bubble(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort_bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
