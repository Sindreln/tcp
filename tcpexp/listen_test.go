package tcpexp

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bundet til %q", listener.Addr())
}
