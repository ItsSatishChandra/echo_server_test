package tcp_server

import (
	"fmt"
	"log"
	"net"

	internal "github.com/ItsSatishChandra/echo_server_test/cmd/internal"
)

func TcpServer(port int) {
	source := "TCP"
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		internal.EchoLogger(source, fmt.Sprintf("Failed to listen on port %v: %v", port, err))
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			internal.EchoLogger(source, fmt.Sprintf("Failed to accept connection: %v", err))
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go internal.ConnectionListener(source, conn)
	}
}
