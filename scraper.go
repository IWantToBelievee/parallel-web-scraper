package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Scraper struct {
	urls []string
	wg   sync.WaitGroup
}

func NewScraper(urls []string) *Scraper {
	return &Scraper{urls: urls}
}

func (s *Scraper) Run() {
	for i, url := range s.urls {
		s.wg.Add(i)
		go s.HTTPRequest(&url)
	}

	s.wg.Wait()

	fmt.Print("Success! \n")
}

func (s *Scraper) HTTPRequest(url *string) {
	defer s.wg.Done()

	response, err := http.Get(*url)
	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s.\n", err)
	}

	fmt.Printf("Done. %d\n", len(body))
}
