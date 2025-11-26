package echo

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ItsSatishChandra/echo_server_test/cmd/internal"
	tcp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/tcp"
	udp_server "github.com/ItsSatishChandra/echo_server_test/cmd/server/udp"
)

func EchoServer(tcpPort int, udpPort int, logFileLocation string) {
	internal.SetLogFileLocation(logFileLocation)
	ctx, cancel := createSignalChannel()
	defer cancel()
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		tcp_server.TcpServer(ctx, tcpPort)
	}()

	go func() {
		defer waitGroup.Done()
		udp_server.UdpServer(ctx, udpPort)
	}()

	waitGroup.Wait()
}

func createSignalChannel() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		internal.EchoLogger("Shutdown Signal Received: ", sig.String())
		cancel()
	}()

	return ctx, cancel
}
