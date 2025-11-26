package tcp_server

import (
	"context"
	"fmt"
	"log"
	"net"

	internal "github.com/ItsSatishChandra/echo_server_test/cmd/internal"
)

func TcpServer(ctx context.Context, port int) {
	source := "TCP"
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		internal.EchoLogger(source, fmt.Sprintf("Failed to listen on port %v: %v", port, err))
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	defer listen.Close()

	acceptChannel := make(chan net.Conn)
	errorChannel := make(chan error)

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				errorChannel <- err
				continue
			}
			acceptChannel <- conn
		}
	}()

	for {
		select {
		case <-ctx.Done():
			internal.EchoLogger(source, "Shutting down TCP server")
			return
		case err := <-errorChannel:
			internal.EchoLogger(source, fmt.Sprintf("Failed to accept connection: %v", err))
		case conn := <-acceptChannel:
			go internal.ConnectionListener(ctx, source, conn)
		}
	}
}
