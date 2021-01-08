package crawler

import (
	"fmt"
	"pearlpig/web_price/shop"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Item struct {
	Node *html.Node `json:"-"`
	Info ItemInfos  `json:"info,omitempty"`
}

func GetItemList(s string, n *html.Node, channel chan<- Item, wg *sync.WaitGroup) {
	crawler := shop.NewListCrawler(s, n)
	fmt.Println("get item num: ", len(crawler.GetItem()))
	for _, v := range crawler.GetItem() {
		wg.Add(1)
		channel <- Item{Node: v, Info: ItemInfos{Server: s}}
	}
}

func (c *Crawler) CrawlList() {
	wg := new(sync.WaitGroup)

	channel := make(chan Item, 100)

	go func() {
		for _, v := range c.PagesNode {
			go GetItemList(v.Server, v.Node, channel, wg)
		}
	}()

	go func() {
		for {
			v, open := <-channel
			if !(open == true) {
				break
			}
			c.Items = append(c.Items, v)
			wg.Done()
		}
	}()
	time.Sleep(time.Microsecond * 1000)
	wg.Wait()
	fmt.Println("len item slice", len(c.Items))
}
