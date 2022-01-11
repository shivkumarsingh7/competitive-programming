package main

import (
	"fmt"
)

// stack
// push
// pop

func main() {
	genSlice := []interface{}{1, 2, 3, 4, 5, "hello", 2.5, 6.7, 4.5, "world", "sample"}
	// ch := make(chan interface{})
	op := &Operation{}
	for _, v := range genSlice {
		go op.push(v)
	}

	fmt.Println(op)
	op.pop("int")
	fmt.Println(op)
	op.pop("float64")
	fmt.Println(op)
	op.pop("string")
	fmt.Println(op)
}

type Ops interface {
	push()
	pop()
}

// type DATATYPE interface{}

type Operation struct {
	// Random DATATYPE
	StringSlice []string
	FloatSlice  []float64
	IntSlice    []int
}

func (o *Operation) push(data interface{}) {
	// fmt.Println("Data: ", data, " Type: ", reflect.TypeOf(data))
	switch data.(type) {
	case int:
		o.IntSlice = append(o.IntSlice, data.(int))
	case float64:
		o.FloatSlice = append(o.FloatSlice, data.(float64))
	case string:
		o.StringSlice = append(o.StringSlice, data.(string))
	}
}

func (o *Operation) pop(data string) {
	switch data {
	case "int":
		o.IntSlice = o.IntSlice[0 : len(o.IntSlice)-1]
	case "float64":
		o.FloatSlice = o.FloatSlice[0 : len(o.FloatSlice)-1]
	case "string":
		o.StringSlice = o.StringSlice[0 : len(o.StringSlice)-1]
	}
}
