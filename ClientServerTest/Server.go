package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)
var letters =[10]string{"one","two","three","four","five","six","seven","eight","nine","ten"}
var message string
func main() {
	fmt.Println("Launching server...")

	listener, err := net.Listen("tcp", ":4547")

	if err != nil {
		fmt.Println(err)
		return
	}
    conn, err := listener.Accept()

	for {
		rand.Seed(time.Now().UTC().UnixNano())
		message=randSeq(rand.Intn(10))
		conn.Write([]byte(message + "\n"))
		time.Sleep(1 * time.Second)
	}

	}

func randSeq(n int) string {
	i:=rand.Intn(len(letters))
	return string(letters[i])
	}

