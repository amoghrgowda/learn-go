package main

import (
	"fmt"
)

func main() {
	random_number()
	sum, text := demo_func(10, 20, "hello ", "world")
	fmt.Println("Sum = ", sum, "\ntext = ", text)
	fmt.Printf("%v is returned of type %T", conv(), conv())
}
