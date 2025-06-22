package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <host> <port> <proxyhost>\n", os.Args[0])
		os.Exit(1)
	}

	targetHost := os.Args[1]
	targetPort := os.Args[2]
	proxyHost := os.Args[3]

	targetAddr := net.JoinHostPort(targetHost, targetPort)

	conn, err := net.DialTimeout("tcp", targetAddr, 3*time.Second)
	if err == nil {
		// Direct connection successful
		handleDirectProxy(conn)
	} else {
		// Direct connection failed, use SSH tunnel
		handleSshTunnelProxy(targetHost, targetPort, proxyHost)
	}
}

func handleDirectProxy(conn net.Conn) {
	defer conn.Close()
	done := make(chan struct{})

	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	<-done
	<-done
}

func handleSshTunnelProxy(host, port, proxyHost string) {
	cmd := exec.Command("ssh", "-W", fmt.Sprintf("%s:%s", host, port), proxyHost)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ssh command failed: %v\n", err)
		os.Exit(1)
	}
}
