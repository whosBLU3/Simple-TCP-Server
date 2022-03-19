package main

// Simple TCP Server
import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	/*  Listen For Connections
	Accept Connections To Server
	Handle The Connections
	*/
	stream, err := net.Listen("tcp", "0.0.0.0:9999")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stream.Close()

	for {
		con, err := stream.Accept()
		fmt.Fprintf(con, "\nSimple TCP Server In Golang")
		if err != nil {
			fmt.Println(err)
			return
		}

		go handler(con)
	}
}

func handler(con net.Conn) {
	for {
		// Read a string only when we encounter a new line char
		// Read Data from connection stream and into the buffer
		data, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}

}
