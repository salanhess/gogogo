package main

import (
	common_util "gogogo/sort_10/common"
	"reflect"
	"testing"
)

func Test_quick1(t *testing.T) {
	tests := common_util.Tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if QuickSort1(tt.Args.Nums, 0, len(tt.Args.Nums)-1); !reflect.DeepEqual(tt.Args.Nums, tt.Want) {
				t.Errorf("QuickSort1() = %v, want %v", tt.Args.Nums, tt.Want)
			}
		})
	}
}
