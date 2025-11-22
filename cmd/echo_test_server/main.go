package main

import (
	echo "github.com/ItsSatishChandra/echo_server_test/cmd/Echo"
)

func main() {
	echo.SetLogFileLocation("")
	echo.EchoServer(20001)
}
