package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"net/http"

	"egd.go/project/common"
)

// Perform health check with interval until canceled
func RunHealthClient(base_url string, interval time.Duration, ctx context.Context) {
	logPrefix := "[HealthClient]"      // Init logPrefix
	port := common.GetPort()           // Get port from .env file
	ticker := time.NewTicker(interval) // Create ticker
	defer ticker.Stop()
	log.Printf("%v Started", logPrefix)
	for { // Loop health checks
		select {
		case <-ctx.Done(): // Check context cancelation
			log.Printf("%v Shut down: %v", logPrefix, ctx.Err())
			return
		case <-ticker.C: // Check server health
			url := fmt.Sprintf("%v:%d%v", base_url, port, common.HealthRoute)
			resp, err := http.Get(url)
			if err != nil { // Handle error
				log.Printf("%v %v", logPrefix, err)
				continue
			}
			defer resp.Body.Close()
			log.Printf("%v Server health: %d %v", logPrefix, resp.StatusCode, http.StatusText(resp.StatusCode))
		}
	}
}
