package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/IbrahimMohammed47/codecrafters-redis-go/commands"
	"github.com/IbrahimMohammed47/codecrafters-redis-go/resp"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
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
		decodedMsg, err := resp.DecodeBytes(msg)
		var respRes resp.Resp
		if err != nil {
			respRes = resp.NewError(err)
		} else if decodedMsg.Type().String() != "<array>" {
			respRes = resp.NewErrorWithString("input must be a resp array")
		} else {
			respRes = commands.HandleCommand(decodedMsg)
		}
		writer := bufio.NewWriter(con)
		err = resp.Encode(writer, respRes)
		if err != nil {
			respRes = resp.NewError(err)
			resp.Encode(writer, respRes)
		}
		writer.Flush()

	}

}
