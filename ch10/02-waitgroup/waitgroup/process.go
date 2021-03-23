package waitgroup

import (
	"log"
	"sync"
	"time"
)

// Crawl collects responses from a list of urls that are passed in.
// It waits for all requests to complete before returning.
func Crawl(sites []string) ([]int, error) {
	start := time.Now()
	log.Printf("starting crawling")
	wg := &sync.WaitGroup{} // init wait counter

	var resps []int
	cerr := &CrawlError{}
	for _, v := range sites {
		wg.Add(1)           // increase wait counter with each site
		go func(v string) { // dispatch url client for each site
			defer wg.Done() // decrease wait counter when processed
			resp, err := GetURL(v)
			if err != nil {
				cerr.Add(err)
				return
			}
			resps = append(resps, resp.StatusCode)
		}(v)
	}
	wg.Wait()
	// we encountered a crawl error
	if cerr.Present() {
		return nil, cerr
	}
	log.Printf("completed crawling in %s", time.Since(start))
	return resps, nil
}
