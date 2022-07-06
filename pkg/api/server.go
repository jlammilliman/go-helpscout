package api

import (
	"os"
)

type Server struct {
	Concurrency int

	Client *Client

	Debug bool

	shutdown       chan struct{}
	doneProcessing chan struct{}
	SigCh          chan os.Signal
	done           chan struct{}
}

// Start starts the request loop for receiving messages and a configurable
// number of Handler routines for message processing. Each process is started in its own goroutine.
// This increases throughput when one of the processess is making an IO call, call to AWS or the
// Alert Hub API.
func (s *Server) Start() {
	go s.startReceiver()
}

// Starts the receiver that handles the receiving messages from the Processor Queue
func (s *Server) startReceiver() {
	return
}
