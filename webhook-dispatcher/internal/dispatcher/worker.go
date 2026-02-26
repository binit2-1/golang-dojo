package dispatcher

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func DispatchWebhooks(urls []string, jsonPayload []byte) {
	var wg sync.WaitGroup

	jobs := make(chan string, len(urls))

	for id := range 5 {
		wg.Add(1)
		go Worker(id, jobs, jsonPayload, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	//channel -> tell no more jobs
	close(jobs)

	//wait all workers to finish
	wg.Wait()

}

func Worker(id int, jobs <-chan string, payload []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Printf("[Worker %d] Failed to send to %s: %v\n", id, url, err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("[Worker %d] Successfully notified %s\n", id, url)
		resp.Body.Close()
	}
}
