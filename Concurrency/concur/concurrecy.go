package concur

type WebsiteChecker func(url string) bool

type Result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	channel := make(chan Result)

	for _, url := range urls {
		go func(url string) {
			channel <- Result{
				url,
				wc(url),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resp := <-channel
		result[resp.string] = resp.bool
	}
	return result
}
