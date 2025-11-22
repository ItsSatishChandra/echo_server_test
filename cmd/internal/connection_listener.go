package internal

import (
	"bufio"
	"net"
)

func ConnectionListener(source string, connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		output := Echo(source, message)
		connection.Write([]byte(output))
	}
}
