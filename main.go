package main

import "reflect"
import "fmt"

type adder func(int, int) int

func makeAdder() adder {
	var (
		a adder
		t = reflect.TypeOf(a)
	)

	add := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		x := args[0].Interface().(int)
		y := args[1].Interface().(int)

		sum := x + y
		return []reflect.Value{reflect.ValueOf(sum)}
	})

	return add.Interface().(adder)
}

func main() {
	add := makeAdder()
	fmt.Printf("the sum of 1 and 2 is %v\n", add(1, 2))
}
