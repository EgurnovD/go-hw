package server

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"golang.org/x/time/rate"

	"egd.go/project/common"
	"github.com/gin-gonic/gin"
)

var stats common.Counters // Statistics for the report

// Return OK
func healthHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Return randomized statuses
func handleHandler(c *gin.Context) {
	pos_status := []int{http.StatusOK, http.StatusAccepted}
	neg_status := []int{http.StatusBadRequest, http.StatusInternalServerError}
	status := pos_status[rand.IntN(len(pos_status))] // Randomize positive status
	if rand.IntN(100) < 30 {                         // Randomize returned status pos/neg
		status = neg_status[rand.IntN(len(neg_status))] // Randomize negative status
	}
	stats.Inc(status) // Increment counter
	c.Status(status)
}

// Return statistical report
func reportHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stats.Report())
}

// Simple rate limiter implementation
func rateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(5, 5)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
		}
	}
}

// Run project server
func StartServer() {
	logPrefix := "[Server]"
	stats = *common.NewCounters() // Init stats
	port := common.GetPort()      // Get port from .env file
	// Setup server
	route := gin.New()
	route.Use(gin.Recovery())
	// route.Use(gin.Logger())
	route.Use(rateLimiter())
	// Configure routes ans handlers
	route.GET(common.HealthRoute, healthHandler)
	route.GET(common.ReportRoute, reportHandler)
	route.POST(common.HandleRoute, handleHandler)
	// Start server
	log.Printf("%v Starting server at :%d", logPrefix, port)
	route.Run(fmt.Sprintf(":%d", port))
}
