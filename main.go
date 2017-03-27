package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn, buffer chan []byte) {
	r := bufio.NewReader(conn)
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("read failed: %v", err)
			}
			_ = conn.Close()
			return
		}

		buffer <- data
	}
}

func stdoutWriter(buffer chan []byte) {
	for {
		_, err := os.Stdout.Write(<-buffer)
		if err != nil {
			log.Printf("write failed: %v", err)
			continue
		}
	}
}

func main() {
	addr := flag.String("addr", "127.0.0.1:1313", "listen address")
	bufLen := flag.Int("buffer", 512, "number of lines to buffer")
	flag.Parse()

	log.SetOutput(os.Stderr)

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make(chan []byte, *bufLen)

	go stdoutWriter(buffer)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept failed: %v", err)
			continue
		}
		go handleConnection(conn, buffer)
	}
}
