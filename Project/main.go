package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"egd.go/project/client"
	"egd.go/project/common"
	"egd.go/project/server"
)

// Log statistical report from server
func logServerReport(base_url string, port int, logPrefix string) {
	url := fmt.Sprintf("%v:%d%v", base_url, port, common.ReportRoute)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%v %v", logPrefix, err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("%v Server report: %v", logPrefix, string(body))
}

func main() {
	log.SetPrefix("[EGD] ")
	logPrefix := "[Main]"
	base_url := "http://0.0.0.0"
	log.Printf("%v Started", logPrefix)
	// Start Server
	go server.StartServer()
	// Run HealthClient
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Graceful shutdown for HealthClient
	go client.RunHealthClient(base_url, time.Second, ctx)
	// Run HandleClients
	var wg sync.WaitGroup
	n_client := 2  // # of HandleClients
	n_workers := 2 // # of workers per client
	n_req := 100   // # of requests per client. Dividable by n_workers * n_job_req
	n_job_req := 5 // # of requests per worker job
	rps := 5       // Requests Per Second for each client
	for i := 0; i < n_client; i++ {
		wg.Add(1)
		go client.RunHandleClient(i, base_url, n_workers, n_req, n_job_req, rps, &wg)
	}
	wg.Wait() // Wait for HandleClients to finish
	logServerReport(base_url, common.GetPort(), logPrefix)
	log.Printf("%v Finished", logPrefix)
}
