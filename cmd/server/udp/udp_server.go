package udp_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
)

func UdpServer(ctx context.Context, port int) {
	source := "UDP"
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		internal.EchoLogger(source, fmt.Sprintf("Failed to listen on port %v: %v", port, err))
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

	for {
		select {
		case <-ctx.Done():
			internal.EchoLogger(source, "Shutting down UDP server")
			return
		default:
		}
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
			continue
		}

		if err != nil {
			internal.EchoLogger(source, fmt.Sprintf("Failed to read from UDP: %v", err))
			continue
		}

		input := string(buffer[:n])
		if input == "" {
			continue
		}

		output := internal.Echo(source, input)

		_, err = conn.WriteToUDP([]byte(output), clientAddr)
		if err != nil {
			select {
			case <-ctx.Done():
				internal.EchoLogger(source, "Shutting down UDP server")
				return
			default:
				if ne, ok := err.(net.Error); ok && ne.Timeout() {
					continue
				}
				internal.EchoLogger(source, fmt.Sprintf("Failed to write to UDP: %v", err))
				continue
			}
		}
	}

}
