package main

import "flag"

func main() {
	listenAddr := flag.String("listen-addr", ":8080", "TCP address to listen on")
	flag.Parse()

	server.Start(&listenAddr)
}
