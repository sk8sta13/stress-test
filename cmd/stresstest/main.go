package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
	wg          sync.WaitGroup
	requesterr  uint32
	status1     uint32
	status2     uint32
	status3     uint32
	status4     uint32
	status5     uint32
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Aplicativo de exemplo com flag --nome",
		Run: func(cmd *cobra.Command, args []string) {
			start := time.Now()
			requestsControl := make(chan struct{}, concurrency)

			for i := 1; i <= requests; i++ {
				wg.Add(1)
				requestsControl <- struct{}{}
				go get(url, requestsControl)
			}
			wg.Wait()

			elapsed := time.Since(start)
			totalRequests := requests - int(requesterr)
			println("\n--------------------------------------------------\n")
			println(fmt.Sprintf("Total execution time..........: %fs", elapsed.Seconds()))
			println(fmt.Sprintf("Total requests made...........: %d", totalRequests))
			println(fmt.Sprintf("Total requests with status 1xx: %d", status1))
			println(fmt.Sprintf("Total requests with status 2xx: %d", status2))
			println(fmt.Sprintf("Total requests with status 3xx: %d", status3))
			println(fmt.Sprintf("Total requests with status 4xx: %d", status4))
			println(fmt.Sprintf("Total requests with status 5xx: %d", status5))
			println("\n--------------------------------------------------\n")
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the service to be tested")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Total number of requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Number of simultaneous calls")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func get(url string, requestsControl <-chan struct{}) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		atomic.AddUint32(&requesterr, 1)
		fmt.Println("Error making request:", err)
		<-requestsControl
	}

	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= 100 && resp.StatusCode < 200:
		atomic.AddUint32(&status1, 1)
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		atomic.AddUint32(&status2, 1)
	case resp.StatusCode >= 300 && resp.StatusCode < 400:
		atomic.AddUint32(&status3, 1)
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		atomic.AddUint32(&status4, 1)
	case resp.StatusCode >= 500:
		atomic.AddUint32(&status5, 1)
	}

	println("The request", url, "returned", resp.StatusCode)
	<-requestsControl
}
