package client

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"net/http"

	"egd.go/project/common"
)

// Perorm n_req POST requests
func workerJob(job_id int, n_req int, rps int, url string, logPrefix string, stats *common.Counters) int {
	logPrefix += fmt.Sprintf("[j %d]", job_id)             // Append logPrefix
	req_count := 0                                         // Init request counter
	interval := time.Duration(1000/rps) * time.Millisecond // Compute interval
	ticker := time.NewTicker(interval)                     // Create ticker
	defer ticker.Stop()
	for range ticker.C { // Loop POST request
		// Do POST request
		resp, err := http.Post(url, "applapplication/json", strings.NewReader(""))
		if err != nil { // Handle error
			log.Printf("%v %v", logPrefix, err)
			continue
		}
		defer resp.Body.Close()
		log.Printf("%v POST: %d %v", logPrefix, resp.StatusCode, http.StatusText(resp.StatusCode))
		// Count request and status
		stats.Inc(resp.StatusCode)
		req_count++
		if req_count >= n_req {
			break
		}
	}
	return req_count // Return number of requests made in the job
}

// Split n_req requests in jobs of n_job_req and run them with lock
func runHandleWorker(worker_id int, url string, n_req int, n_job_req int, rps int,
	logPrefix string, stats *common.Counters, worker_wg *sync.WaitGroup, mx *sync.Mutex) {
	defer worker_wg.Done()
	logPrefix += fmt.Sprintf("[w %d]", worker_id) // Append logPrefix
	req_count := 0                                // Init request counter
	job_count := 0                                // Init job counter
	for req_count < n_req {                       // Loop jobs
		mx.Lock() // Acquire exclusive lock
		req_count += workerJob(job_count, n_job_req, rps, url, logPrefix, stats)
		job_count++
		mx.Unlock()                        // Release the lock
		time.Sleep(500 * time.Millisecond) // Let other workers lock
	}
}

// Perform n_req request to base_url with n_worker workers in groups of n_job_req
func RunHandleClient(client_id int, base_url string, n_worker int, n_req int, n_job_req int, rps int, wg *sync.WaitGroup) {
	defer wg.Done()
	logPrefix := fmt.Sprintf("[HadleClient %d]", client_id)           // Init logPrefix
	stats := *common.NewCounters()                                    // Init stats
	port := common.GetPort()                                          // Get port from .env file
	url := fmt.Sprintf("%v:%d%v", base_url, port, common.HandleRoute) // Compose url
	log.Printf("%v Started", logPrefix)
	var mx sync.Mutex               // Mutex for worker rotation
	var worker_wg sync.WaitGroup    // WaitGroup for worker finish
	for i := 0; i < n_worker; i++ { // Create workers
		worker_wg.Add(1)
		go runHandleWorker(i, url, n_req/n_worker, n_job_req, rps, logPrefix, &stats, &worker_wg, &mx)
	}
	worker_wg.Wait() // Wait for workers
	// Serialize client report to pretty JSON
	jsonData, err := json.MarshalIndent(stats.Report(), "", "    ")
	if err != nil { // Handle error
		log.Printf("%v Error marshalling JSON: %v", logPrefix, err)
	}
	log.Printf("%v Finished. Report: %v", logPrefix, string(jsonData))
}
