package main

import (
	"fmt"
)

func main() {
	fmt.Println("Факториал числа", factorial(4));
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	//изменение в комментарии отправить на github
	return n * factorial(n - 1)
}
