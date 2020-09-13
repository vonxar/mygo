package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	v   map[string]bool
	mux sync.Mutex
}

//urlに一度でもマグっったことがあるならば,tureを返す
//巡ってないなら、falseを返してmap@[url]=trueとする
func (c *SafeCache) isFindAndSet(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.v[url]
	if ok && val {
		return true
	}
	c.v[url] = true
	return false

}

type Fetcher interface {
	//FetchはURLの本文を返し、
	// そのページで見つかったURLのスライス。
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	cache := SafeCache{v: make(map[string]bool)}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go crawl(url, depth, fetcher, &cache, wg)
	wg.Wait()
}

//クロールはフェッチャーを使用して再帰的にクロールします
// urlで始まり、最大の深さまでのページ。
func crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	// TODO：並行してURLを取得します。
	// TODO：同じURLを2回フェッチしないでください。
	//この実装はどちらも行いません：
	defer wg.Done()

	if depth <= 0 {
		return
	}
	if cache.isFindAndSet(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawl(u, depth-1, fetcher, cache, wg)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcherは、既定の結果を返すFetcherです。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcherは入力済みのfakeFetcherです。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
