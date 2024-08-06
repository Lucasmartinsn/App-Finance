package main

import (
	"development/application/fiance/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
