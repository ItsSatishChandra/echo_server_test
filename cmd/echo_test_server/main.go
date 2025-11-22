package main

import (
	"flag"

	echo_server "github.com/ItsSatishChandra/echo_server_test/cmd/Echo"
)

func main() {

	tcpPort := flag.Int("tcp-port", 20001, "Port for TCP Echo Server")
	udpPort := flag.Int("udp-port", 20002, "Port for UDP Echo Server")
	logFileLocation := flag.String("log-file", "", "Log file location")
	flag.Parse()

	echo_server.EchoServer(*tcpPort, *udpPort, *logFileLocation)
}
