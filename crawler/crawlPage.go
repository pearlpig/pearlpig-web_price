package crawler

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"pearlpig/web_price/shop"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Nodes []*html.Node

var itemNode Nodes

type Page struct {
	Server string
	Target string
	Node   *html.Node
}

const (
	maxIdleConns        int = 0
	maxIdleConnsPerHost int = 1000
	timeout             int = 180
)

func Random(i int) int {
	return (i * i) / 700
}
func (crawler *Crawler) CrawlPage() {
	allPageNum := cap(crawler.PagesNode)
	channelReq := make(chan Page, 200)
	channelRes := make(chan Page, 50)
	wg := new(sync.WaitGroup)
	wg.Add(allPageNum)
	client := createHttpClient()
	// delayUnitTime := time.Millisecond
	go func() {
		for {
			v, open := <-channelRes
			if !(open == true) {
				fmt.Println("finish respones")
				break
			}
			crawler.PagesNode = append(crawler.PagesNode, v)
			wg.Done()
		}

	}()
	go func() {
		for i := 0; i < allPageNum; i++ {
			// delayTime := delayUnitTime * time.Duration(GenerateRangeNum())
			// fmt.Println(delayTime)
			// share the same client
			go Crawl(channelReq, channelRes, wg, client, 0)
		}
	}()
	cnt := 0
	for _, s := range crawler.ReqServer {
		for page := 0; page < s.PageNum; page++ {
			cnt++
			p := Page{Server: s.Name, Target: s.Url.Home + s.Url.Search + crawler.ReqItem + s.Url.Page + strconv.Itoa(page) + s.Url.Sort}
			fmt.Println("crawling...", p)
			channelReq <- p
		}
	}
	close(channelReq)

	wg.Wait()
	if !(len(crawler.PagesNode) == allPageNum) {
		panic("Wrong page node number")
	}
}

func Crawl(channelReq chan Page, channelRes chan Page, wg *sync.WaitGroup, client *http.Client, d time.Duration) {

	p := shop.NewPageParser(client)
	for {
		v, open := <-channelReq
		if !(open == true) {
			//finish request
			break
		}
		p.SetTarget(v.Target)
		// fmt.Println("delay ", d)
		time.Sleep(d)
		v.Node = p.Parse()
		channelRes <- v
	}
}

func createHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
		},
		Timeout: time.Duration(timeout) * time.Second,
	}
	return client
}

func GenerateRangeNum() int64 {
	randNum, _ := rand.Int(rand.Reader, big.NewInt(500))

	return randNum.Int64()
}
