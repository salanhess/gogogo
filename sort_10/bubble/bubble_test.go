package main

import (
	"gogogo/sort_10/common"
	"reflect"
	"testing"
)

func Test_sort1(t *testing.T) {
	tests := common_util.Tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := sort_bubble(tt.Args.Nums); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("sort_select() = %v, want %v", got, tt.Want)
			}
		})
	}
}

func Test_sort2(t *testing.T) {
	tests := common_util.Tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := sort_bubble2(tt.Args.Nums); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("sort_select() = %v, want %v", got, tt.Want)
			}
		})
	}
}

func Test_sort3(t *testing.T) {
	tests := common_util.Tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := sort_bubble3(tt.Args.Nums); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("sort_select() = %v, want %v", got, tt.Want)
			}
		})
	}
}
