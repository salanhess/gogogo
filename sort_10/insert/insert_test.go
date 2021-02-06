package main

import (
	"gogogo/sort_10/common"
	"reflect"
	"testing"
)

func Test_sort_Insert(t *testing.T) {
	tests := common_util.Tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := sort_insert(tt.Args.Nums); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("sort_select() = %v, want %v", got, tt.Want)
			}
		})
	}
}
