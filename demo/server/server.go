package main

import (
	"fenginx/fnet"
)

func main() {
	s := fnet.NewServer("test server")
	s.Serve()
}
