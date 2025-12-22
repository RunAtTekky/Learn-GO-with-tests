package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/runattekky/go-specs-greet/adapters"
	"github.com/runattekky/go-specs-greet/adapters/httpserver"
	"github.com/runattekky/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port    = "8080"
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, driver)
}
