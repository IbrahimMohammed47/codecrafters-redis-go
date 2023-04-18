package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	// Uncomment this block to pass the first stage
	// "net"
	// "os"
	"github.com/IbrahimMohammed47/codecrafters-redis-go/resp"
)

func main() {

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	fmt.Println("Redis replica is listening on port 6379 ...")

	for {
		con, err := l.Accept()
		fmt.Println("A new connection from " + con.LocalAddr().String())
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(con)
	}

}

func handleConnection(con net.Conn) {
	// var msg []byte
	for {
		msg := make([]byte, 1024)
		_, err := con.Read(msg)
		if err != nil {
			// fmt.Println("Error reading msg: ", err.Error())
			break
		}
		res := "PONG"
		respRes := resp.NewString(res)
		writer := bufio.NewWriter(con)
		resp.Encode(writer, respRes)
		writer.Flush()

	}

}
