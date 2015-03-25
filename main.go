package main

import (
	"github.com/wfernandes/golife/server"
)

func main() {
	server := server.NewTimelineServer()
	server.Start()
}
