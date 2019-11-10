package main

import (
	"flag"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/infrastructure/server"
)

func main() {
	listenAddr := flag.String("listen-addr", ":8080", "TCP address to listen on")
	flag.Parse()

	server.Start(*listenAddr)
}
