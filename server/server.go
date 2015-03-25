package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	mx := mux.NewRouter()

	mx.HandleFunc("/", s.ServeHome)
	mx.HandleFunc("/timeline", s.ServeJSON)

	log.Print("Starting Timeline Server...")
	err := http.ListenAndServe(":8080", mx)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
		return
	}

}

func NewTimelineServer() *TimelineServer {
	return &TimelineServer{}
}
