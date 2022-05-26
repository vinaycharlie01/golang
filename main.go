package main

import "fmt"

func main() {
	// fmt.Println("enter your number:")
	// var a float64
	// fmt.Scanln(&a)
	// fmt.Println("enter your number:")
	// var b float64
	// fmt.Scanln(&b)
	// b = math.Min(a, b)
	// if b == a {
	// 	d := b / a
	// 	fmt.Println(d)

	// } else {
	// 	e := a / b
	// 	fmt.Println(e)
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz", i)
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz", i)
	// 	}
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz", i)
	// 	}
	// }

	b := hello("vinay")
	fmt.Println(b)
}

func hello(x interface{}) (b []interface{}) {

	switch x.(type) {
	case string:
		b = append(b, true)
		return b
	case float64:
		b = append(b, false)
		return b
	}
	return append(b, b)
}
