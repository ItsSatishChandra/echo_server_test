package echo

import (
	"bufio"
	"net"
)

func ConnectionListener(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		output := Echo(message)
		connection.Write([]byte(output))
	}
}
