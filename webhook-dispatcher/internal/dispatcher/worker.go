package dispatcher

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func DispatchWebhooks(urls []string, jsonPayload []byte) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(targetURL string) {
			defer wg.Done()
			resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(jsonPayload))
			if err != nil {
				fmt.Printf("Failed to send to %s\n", targetURL)
				return
			}

			defer resp.Body.Close()

		}(url)
	}
	wg.Wait()

}
