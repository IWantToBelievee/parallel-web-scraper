package main

func main() {
	urls := []string{"https://google.com", "https://golang.org", "https://github.com"}

	scraper := NewScraper(urls)
	scraper.Run()
}
