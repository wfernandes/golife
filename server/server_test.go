package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/wfernandes/golife/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {

	Context("handlers", func() {
		It("should return a 200 response with helloworld text", func() {

			s := server.NewTimelineServer()
			testServer := httptest.NewServer(http.HandlerFunc(s.ServeHome))
			defer testServer.Close()
			resp, _ := http.Get(testServer.URL)
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
			Expect(ioutil.ReadAll(resp.Body)).To(ContainSubstring("helloworld"))
		})

		It("should return JSON formatted hello world from /json endpoint", func() {
			s := server.NewTimelineServer()
			testServer := httptest.NewServer(http.HandlerFunc(s.ServeJSON))
			defer testServer.Close()

			resp, _ := http.Get(fmt.Sprintf("%s/json", testServer.URL))
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
			Expect(ioutil.ReadAll(resp.Body)).To(ContainSubstring("{\"hello\":\"world\"}"))
		})
	})
})
