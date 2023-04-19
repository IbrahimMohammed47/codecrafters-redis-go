package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println(1, err)
		// handle error
	}
	defer conn.Close()
	msg := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	conn.Write([]byte(msg))

	rcvdMsg := make([]byte, 1024)
	conn.Read(rcvdMsg)
	fmt.Println(string(rcvdMsg))

	// msg2 := "*1\r\n$4\r\nPING\r\n"
	// msg2 := "PING"
	msg2 := "*1\r\n$4\r\nPING\r\n"
	conn.Write([]byte(msg2))

	rcvdMsg2 := make([]byte, 1024)
	conn.Read(rcvdMsg2)
	fmt.Println(string(rcvdMsg2))

	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println(2, err)
	// 	// handle error
	// }
	// fmt.Println(status)

}
