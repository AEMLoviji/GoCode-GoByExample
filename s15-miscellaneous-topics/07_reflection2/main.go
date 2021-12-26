package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// testCase1()
	// testCase2()
	// testCase3()
	testCase4()
}

func testCase1() {
	t := reflect.TypeOf(10)
	fmt.Println(t) // int

	fmt.Println(reflect.TypeOf(10.31))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf([]int{1, 2}))
	fmt.Println(reflect.TypeOf([...]int{1, 2}))

	var f io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(f))
}

func testCase2() {
	i := 10
	iV := reflect.ValueOf(i)
	fmt.Println(iV)
	fmt.Printf("%v\n", iV)
	fmt.Println(iV.String())
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(i).String())

	t := iV.Type()
	fmt.Println(t.String())
}

func testCase3() {

	f := 10.32

	fV := reflect.ValueOf(f)
	fT := reflect.TypeOf(f)

	fmt.Printf("fV=%v, fV.Kind()=%v, fV.Type()=%v\n",
		fV, fV.Kind(), fV.Type())
	//  10.32 float64  float64

	fmt.Printf("fV.String()=%v fv.Float()=%v fV.Interface()=%v fV.Interface().(float64)=%v\n",
		fV.String(), fV.Float(), fV.Interface(), fV.Interface().(float64))
	// <float64 Value> 10.32     10.32           10.32

	fmt.Printf("fT=%v, fT.Kind()=%v fT.String()=%v \n",
		fT, fT.Kind(), fT.String())
	// float64 float64 float64

	if (fV.Kind() == reflect.Float64) && (fT.Kind() == reflect.Float64) {
		fmt.Println("float64 is underlying type.")
	}
}

func testCase4() {

	type myFloat float64
	var f myFloat = 10.31
	fV := reflect.ValueOf(f)
	fT := reflect.TypeOf(f)

	fmt.Printf("fV=%v, fV.Kind()=%v, fV.Type()=%v\n",
		fV, fV.Kind(), fV.Type())
	//  10.31 float64  main.myFloat

	fmt.Printf("fV.String()=%v fv.Float()=%v fV.Interface()=%v fV.Interface().(myFloat)=%v\n",
		fV.String(), fV.Float(), fV.Interface(), fV.Interface().(myFloat))
	// <main.myFloat Value> 10.31     10.31           10.31

	fmt.Printf("fT=%v, fT.Kind()=%v fT.String()=%v \n",
		fT, fT.Kind(), fT.String())
	// main.myFloat float64 main.myFloat

	if (fV.Kind() == reflect.Float64) && (fT.Kind() == reflect.Float64) {
		fmt.Println("The King is always the underlying type.")
	}
}
