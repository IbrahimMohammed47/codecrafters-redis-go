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

	msg2 := "*1\r\n$4\r\nping\r\n"
	conn.Write([]byte(msg2))

	rcvdMsg2 := make([]byte, 1024)
	conn.Read(rcvdMsg2)
	fmt.Println(string(rcvdMsg2))

	// msg3 := "*3\r\n$3\r\nSET\r\n+num\r\n:7\r\n"
	// conn.Write([]byte(msg3))

	// rcvdMsg3 := make([]byte, 1024)
	// conn.Read(rcvdMsg3)
	// fmt.Println(string(rcvdMsg3))

	// msg3 := "*5\r\n$3\r\nSET\r\n+num\r\n:7\r\n$2\r\nPX\r\n$3\r\n100\r\n"
	// conn.Write([]byte(msg3))

	// rcvdMsg3 := make([]byte, 1024)
	// conn.Read(rcvdMsg3)
	// fmt.Println(string(rcvdMsg3))

	// msg4 := "*2\r\n$3\r\nGET\r\n+num\r\n"
	// conn.Write([]byte(msg4))

	// rcvdMsg4 := make([]byte, 1024)
	// conn.Read(rcvdMsg4)
	// fmt.Println(string(rcvdMsg4))

	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println(2, err)
	// 	// handle error
	// }
	// fmt.Println(status)

}
