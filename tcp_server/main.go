package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go func(c net.Conn) {
			err := c.SetDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				fmt.Println("Connection timed out")
			}
			request(c)
			response(c)

			defer c.Close()
		}(conn)
	}
}

func request(c net.Conn) {
	i := 0
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			method := strings.Fields(ln)[0]
			fmt.Println("***METHOD", method)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(c net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}
