package echo

import (
	"sync"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
	tcp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/tcp"
	udp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/udp"
)

func EchoServer(tcpPort int, udpPort int, logFileLocation string) {
	internal.SetLogFileLocation("")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		tcp_server.TcpServer(tcpPort)
	}()

	go func() {
		defer waitGroup.Done()
		udp_server.UdpServer(udpPort)
	}()

	waitGroup.Wait()
}
