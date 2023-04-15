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
	msg := "PING"
	conn.Write([]byte(msg))

	rcvdMsg := make([]byte, 1024)
	conn.Read(rcvdMsg)
	fmt.Println(string(rcvdMsg))

	msg2 := "PING"
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
