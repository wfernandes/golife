package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TimelineServer struct{}

type Timeline struct {
	Data map[string]string
}

func (s *TimelineServer) ServeHome(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "helloworld")
}

func (s *TimelineServer) ServeJSON(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	data["hello"] = "world"
	t := Timeline{
		Data: data,
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
