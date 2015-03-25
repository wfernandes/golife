package server

import (
	"io"
	"log"
	"net/http"
)

type TimelineServer struct{}

func (s *TimelineServer) ServeHome(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "helloworld")
}

func (s *TimelineServer) Start() {
	http.HandleFunc("/", s.ServeHome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}

func NewTimelineServer() *TimelineServer {
	return &TimelineServer{}
}
