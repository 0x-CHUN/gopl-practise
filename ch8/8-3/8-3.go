package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	c := conn.(*net.TCPConn)

	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	_ = c.CloseWrite()
	<-done
	_ = c.CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
