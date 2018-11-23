package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)
var letters =[10]string{"one","two","three","four","five","six","seven","eight","nine","ten"}
var message string

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	sigsServer := make(chan os.Signal, 1)
	doneServer := make(chan bool, 1)
	listener, err := net.Listen("tcp", ":4547")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Launching server...")
go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)

	}
}()
	signal.Notify(sigsServer, syscall.SIGINT, syscall.SIGTERM)
	handleServer(sigsServer, doneServer)
	time.Sleep(100 * time.Second)
	}

func  handleConnection (conn net.Conn)  {
	message = randSeq(rand.Intn(10))
	conn.Write([]byte(message + "\n"))
	time.Sleep(1 * time.Second)}


func randSeq(n int) string {
	i:=rand.Intn(len(letters))
	return string(letters[i])
	}


func handleServer(sigs chan os.Signal,done chan bool) {
	s := <-sigs // ожидание сигнала
	fmt.Println("Программа завершится через 3 секунды. Возникшее событие:", s)
	time.Sleep(3 * time.Second)
	done <- true
	os.Exit(0)
}