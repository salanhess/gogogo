//package _1_exceldriver
package main

import (
	"testing"
)

/*分离测试数据、测试逻辑
明确出错信息
部分错误可以继续
go 语法来实践表格驱动

命令行验证代码覆盖率：先生成c.out 再转成html阅读
go test -coverprofile=c.out
go tool cover -html c.out

性能测试
func BenchmarkSubstr(bt *testing.B)
go test -bench .

性能分析
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
*/

func TestTriangel(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{12, 35, 37},
		//{math.MaxInt32, 1,math.MinInt32}, //溢出变为最小的了
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangel(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriange(%d,%d); "+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
	triangel()
}

func BenchmarkSubstr(bt *testing.B) {
	var a, b, c int = 30000, 40000, 50000
	for i := 0; i < bt.N; i++ {
		if actual := calcTriangel(a, b); actual != c {
			bt.Errorf("calcTriange(%d,%d); "+"got %d; expected %d", a, b, actual, c)
		}
	}
}
