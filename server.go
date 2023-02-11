package main

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,          //1Mb
		ReadTimeout:    10 * time.Second, //10 Sec
		WriteTimeout:   10 * time.Second, //10 Sec
	}

	return s.httpServer.ListenAndServe()
}
