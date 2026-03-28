package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Util() {
	name := "Paras"
	message := "Hello " + name + "!"
	Age := 25
	fmt.Println(message, "Age is: ", Age)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i - 3)
		} else {
			fmt.Println(50)
		}
	}

	fmt.Println(add(5, 100))
}
