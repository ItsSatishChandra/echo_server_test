package udp_server

import (
	"fmt"
	"log"
	"net"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
)

func UdpServer(port int) {
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

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			internal.EchoLogger(source, fmt.Sprintf("Failed to read from UDP: %v", err))
			continue
		}

		input := string(buffer[:n])
		output := internal.Echo(source, input)

		_, err = conn.WriteToUDP([]byte(output), clientAddr)
		if err != nil {
			log.Printf("Failed to write to UDP: %v", err)
			internal.EchoLogger(source, fmt.Sprintf("Failed to write to UDP: %v", err))
			continue
		}
	}

}
