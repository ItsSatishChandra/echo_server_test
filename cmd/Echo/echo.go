package echo

import (
	"sync"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
	tcp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/tcp"
	udp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/udp"
)

func EchoServer(port int) {
	internal.SetLogFileLocation("")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		tcp_server.TcpServer(port)
	}()

	go func() {
		defer waitGroup.Done()
		udp_server.UdpServer(port + 1)
	}()

	waitGroup.Wait()
}
