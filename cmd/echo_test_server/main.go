package main

import (
	echo_server "github.com/ItsSatishChandra/echo_server_test/cmd/echo"
)

func main() {
	echo_server.EchoServer(20001)
}
