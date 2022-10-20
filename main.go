package main

import (
	"launchpad/logging"
	"launchpad/server"
)

func main() {
	logging.NewGlobalLogger()
	s := server.NewServer(3000)
	s.Listen()
}
