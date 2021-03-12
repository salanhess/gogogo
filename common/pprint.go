package common

import (
	"fmt"
	"reflect"
)

func PrintVal(model interface{}) {
	val := reflect.ValueOf(model)
	typ := reflect.Indirect(val).Type()
	fmt.Printf("Model type is %s,val is %v \n", typ, val)
}

func PrintVals(models []interface{}) {
	for _, model := range models {
		val := reflect.ValueOf(model)
		typ := reflect.Indirect(val).Type()
		fmt.Printf("Model type is %s,val is %v \n", typ, val)
	}
}
