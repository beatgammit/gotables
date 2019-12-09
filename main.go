package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var config string

type route struct {
	src  string
	dest string
}

func (r route) handle(src net.Conn) {
	defer src.Close()

	dst, err := net.Dial("tcp", r.dest)
	if err != nil {
		log.Printf("Error connecting to destination: %s\n", r.dest)
		return
	}

	defer dst.Close()

	log.Printf("Routing %s -> %s", r.src, r.dest)

	go io.Copy(src, dst)
	io.Copy(dst, src)

	log.Printf("Connection closed: %s -> %s", r.src, r.dest)
}

func (r route) listen() {
	log.Printf("Creating listener: %s -> %s\n", r.src, r.dest)

	l, err := net.Listen("tcp", r.src)
	if err != nil {
		log.Fatalf("Error setting up listener: %s -> %s: %s", r.src, r.dest, err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
		}
		go r.handle(conn)
	}
}

func init() {
	flag.StringVar(&config, "config", "config.txt", "config file with rules")
	flag.Parse()
}

func readConfig(file string) []route {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable to open file: %s", err)
	}
	defer f.Close()

	var routes []route

	b := bufio.NewReader(f)
	for {
		line, err := b.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading file: %s", err)
		}

		line = strings.TrimSpace(line)

		if line[0] == '#' {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			routes = append(routes, route{parts[0], parts[1]})
		}
	}

	return routes
}

func main() {
	routes := readConfig(config)
	if len(routes) == 0 {
		log.Fatalln("No routes, refusing to start")
	}
	for _, r := range routes {
		go r.listen()
	}
	<-chan bool(nil)
}
