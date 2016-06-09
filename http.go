package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	tcp, _ := net.Listen("tcp", ":8090")
	fmt.Print("build...")
	conn, _ := tcp.Accept()
	fmt.Print("build...end")
	reader := make([]byte, 100)
	httpReq := ""
	for {
		for {
			_, error := conn.Read(reader)
			httpReq = httpReq + string(reader)
			if error == io.EOF {
				conn.Write([]byte("fuck"))
				fmt.Println(httpReq, "fuck")
				break
			}
		}
		tcp.Close()
		break
	}
}
