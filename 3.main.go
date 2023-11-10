package main

import "fmt"

func main() {
	i := 10
	//j := "Hello Cedric"
	//k := "You will succeed"
	//var pointer_integer *int = &i
	//fmt.Println(pointer_integer)
	//var pointer_string = &j
	//fmt.Println(pointer_string)
	//pointer_short_hand := &k
	//fmt.Println(pointer_short_hand)

	var pointer_integer *int = &i
	fmt.Println(pointer_integer)
	var pointer_string = &i
	fmt.Println(pointer_string)
	pointer_short_hand := &i
	fmt.Println(pointer_short_hand)

}
