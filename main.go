package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

// RequestPayload represents the expected JSON payload structure.
type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

// ResponsePayload represents the structure of the JSON response.
type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func main() {
	http.HandleFunc("/process-single", processSingleHandler)
	http.HandleFunc("/process-concurrent", processConcurrentHandler)

	port := 8000
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}

func processSingleHandler(w http.ResponseWriter, r *http.Request) {
	processHandler(w, r, sortSingleArray)
}

func processConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	processHandler(w, r, sortConcurrentArrays)
}

func processHandler(w http.ResponseWriter, r *http.Request, sortFunc func([]int) []int) {
	var requestData RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	var sortedArrays [][]int
	for _, arr := range requestData.ToSort {
		sortedArrays = append(sortedArrays, sortFunc(arr))
	}

	elapsedTime := time.Since(startTime).Nanoseconds()

	response := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       elapsedTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sortSingleArray(arr []int) []int {
	// Sorting a single array sequentially
	sort.Ints(arr)
	return arr
}

func sortConcurrentArrays(arr []int) []int {
	// Sorting a single array concurrently (for demonstration purposes)
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))

	for _, v := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var result []int
	for sortedVal := range ch {
		result = append(result, sortedVal)
	}

	sort.Ints(result)
	return result
}
