package internal

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
)

func ConnectionListener(ctx context.Context, source string, connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		messageChannel := make(chan string)
		errorChannel := make(chan error)

		go func() {
			message, err := reader.ReadString('\n')
			if err != nil {
				errorChannel <- err
				return
			}
			messageChannel <- message
		}()

		select {
		case <-ctx.Done():
			EchoLogger(source, "Closing connection listener due to shutdown signal")
			return
		case err := <-errorChannel:
			if err != nil && !isIgnorableError(err) {
				EchoLogger(source, fmt.Sprintf("Error reading from connection: %v", err.Error()))
				return
			}
		case msg := <-messageChannel:
			output := Echo(source, msg)
			connection.Write([]byte(output))
		}
	}
}

func isIgnorableError(err error) bool {
	if err == io.EOF {
		return true
	}
	// For Windows "forcibly closed by remote host"
	if opErr, ok := err.(*net.OpError); ok {
		if sysErr, ok := opErr.Err.(*os.SyscallError); ok {
			if sysErr.Err.Error() == "WSARecv: An existing connection was forcibly closed by the remote host." {
				return true
			}
		}
	}
	return false
}
