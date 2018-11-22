package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b[] byte) (n int, e error)
}

type Writer interface {
	Write(b[] byte) (n int, e error)
}

func main()  {
	var w Writer
	w=os.Stdout
	fmt.Println(w,"hello word \n")
	fmt.Scanln()
}