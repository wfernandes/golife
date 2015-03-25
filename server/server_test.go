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

		var (
			s          *server.TimelineServer
			testServer *httptest.Server
		)

		BeforeEach(func() {
			s = server.NewTimelineServer()

		})

		AfterEach(func() {
			testServer.Close()
		})

		It("should return a 200 response with helloworld text", func() {

			testServer = httptest.NewServer(http.HandlerFunc(s.ServeHome))
			defer testServer.Close()
			resp, _ := http.Get(testServer.URL)
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
			Expect(ioutil.ReadAll(resp.Body)).To(ContainSubstring("helloworld"))
		})

		It("should return JSON formatted hello world from /timeline endpoint", func() {

			testServer = httptest.NewServer(http.HandlerFunc(s.ServeJSON))
			defer testServer.Close()

			resp, _ := http.Get(fmt.Sprintf("%s/timeline", testServer.URL))
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
			body, _ := ioutil.ReadAll(resp.Body)
			Expect(body).To(ContainSubstring("\"Headline\":\"Hello World!\""))
			Expect(body).To(ContainSubstring("\"StartDate\""))
		})
	})
})
