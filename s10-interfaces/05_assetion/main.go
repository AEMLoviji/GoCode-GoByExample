package main

import "fmt"

func main() {
	var message interface{} = "Hello"

	s, ok := message.(string)
	if ok {
		fmt.Printf("%q %T\n", s, message)
	} else {
		fmt.Printf("value is not a string - %q %T\n", s, message)
	}

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	f := 20.355
	fmt.Printf("%d %T %T \n", int(f), int(f), f)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	var num interface{} = 10
	println(num.(int) + 20)
	fmt.Printf("%d %T \n", num, num)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	//switch i := message.(type) {
	//case int:
	//	fmt.Printf("%T %v", i, i)
	//case float64:
	//	fmt.Printf("%T %v", i, i)
	//case bool, string:
	//	fmt.Println("type is bool or string")
	//case func(int) float64:
	//	fmt.Printf("%T", i)
	//case nil:
	//	fmt.Println("x is nil")
	//default:
	//	fmt.Println("do not know the type")
	//}
}
