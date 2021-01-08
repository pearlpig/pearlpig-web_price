package crawler

import (
	"fmt"
	"pearlpig/web_price/shop"
	"sync"
)

func (c *Crawler) CrawlItem() {
	itemNum := len(c.Items)
	wg := new(sync.WaitGroup)
	wg.Add(itemNum)
	fmt.Println("item number: ", itemNum)
	for i := 0; i < len(c.Items); i++ {
		GetItemInfo(&c.Items[i], wg)
	}
	wg.Wait()
}

func GetItemInfo(i *Item, wg *sync.WaitGroup) {
	defer wg.Done()
	s := shop.NewItemCrawler(i.Info.Server, i.Node)
	// fmt.Println(fmt.Sprintf("name:%s\nurl:%s\nprice:%d\n\n", s.GetName(), s.GetUrl(), s.GetPrice()))
	i.Info.Name = s.GetName()
	i.Info.Url = s.GetUrl()
	price := s.GetPrice()
	i.Info.Price = &price
}
