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
	_, err = conn.Read(rcvdMsg)
	fmt.Println(string(rcvdMsg))
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println(2, err)
	// 	// handle error
	// }
	// fmt.Println(status)

}
