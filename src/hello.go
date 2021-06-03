package main

import (
	"fmt"
	"math"
)

func main() {
	const myConst float64 = math.Sin(1.57)

	fmt.Println("%v, %T\n", myConst, myConst)

	//=> "./hello.go:16:8: const initializer math.Sin(1.57) is not a constant"

}

// var (
// 	itStartsWith     string = "This is a string"
// 	oneThing         string = "This is another string"
// 	iDontKnowWhy     int    = 3
// 	doesntEvenMatter int    = 11
// )
