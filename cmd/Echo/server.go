package echo

import (
	"fmt"
	"log"
	"net"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
)

func EchoServer(port int) {
	internal.SetLogFileLocation("")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go internal.ConnectionListener(conn)
	}
}
