package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type TimelineServer struct{}

type Timeline struct {
	Events []Event
}

type Event struct {
	StartDate time.Time
	Headline  string
}

func (s *TimelineServer) ServeHome(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "helloworld")
}

func (s *TimelineServer) ServeJSON(resp http.ResponseWriter, req *http.Request) {

	e := Event{
		StartDate: time.Now(),
		Headline:  "Hello World!",
	}

	t := Timeline{
		Events: []Event{e},
	}
	output, err := json.Marshal(t)

	if err != nil {
		log.Fatal("Json Marshalling: ", err.Error())
	}

	io.WriteString(resp, string(output))

}

func (s *TimelineServer) Start() {
	http.HandleFunc("/", s.ServeHome)
	http.HandleFunc("/json", s.ServeJSON)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}

}

func NewTimelineServer() *TimelineServer {
	return &TimelineServer{}
}
