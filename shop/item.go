package shop

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type itemInfo struct {
	node        *html.Node
	isNameFind  bool
	isUrlFind   bool
	isPriceFind bool
	name        string
	url         string
	price       int
}

func NewItemInfo(n *html.Node) *itemInfo {
	i := new(itemInfo)

	i.isNameFind = false
	i.isUrlFind = false
	i.isPriceFind = false

	return i
}
func (i *itemInfo) SetName(itemName string) {
	i.name = itemName
}
func (i *itemInfo) GetName() string {
	return i.name
}
func (i *itemInfo) SetUrl(itemUrl string) {
	i.url = itemUrl
}
func (i *itemInfo) GetUrl() string {
	return i.url
}
func (i *itemInfo) SetPrice(itemPrice int) {
	i.price = itemPrice
}
func (i *itemInfo) GetPrice() int {
	return i.price
}

type itemCrawler interface {
	GetName() string
	GetUrl() string
	GetPrice() int
}

// Add new item info crawler method below
func NewItemCrawler(s string, n *html.Node) itemCrawler {
	switch s {
	case "shopee":
		// fmt.Println("create shopee item crawler")
		return NewShopeeItemInfo(n)
	case "yahoo":
		// fmt.Println("create yahoo item crawler")
		return NewYahooItemInfo(n)
	case "ruten":
		// fmt.Println("create ruten item crawler")
		return NewRutenItemInfo(n)
	// Add new item crawler constructor
	default:
		panic("unknown server name")
	}
}

func getContent(n *html.Node) string {
	var buf bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buf.WriteString(c.Data)
	}
	content := strings.ReplaceAll(buf.String(), ",", "")

	return content
}

// shopee info crawler
type shopeeItemInfo struct {
	itemInfo
}

func NewShopeeItemInfo(n *html.Node) *shopeeItemInfo {
	s := new(shopeeItemInfo)

	s.isNameFind = false
	s.isUrlFind = false
	s.isPriceFind = false
	s.FindName(n)
	s.FindUrl(n)
	s.FindPrice(n)
	return s
}
func (s *shopeeItemInfo) FindName(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "_1NoI8_ _16BAGk" {
				itemName := getContent(n)
				s.isNameFind = true
				s.SetName(itemName)
				break
			}
		}
	}
	if !(s.isNameFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindName(c)
		}
	}
}
func (s *shopeeItemInfo) FindUrl(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				s.SetUrl("https://shopee.tw/" + a.Val)
				s.isUrlFind = true
				break
			}
		}
	}
	if !(s.isUrlFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindUrl(c)
		}
	}
}
func (s *shopeeItemInfo) FindPrice(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "_341bF0" {
				tmp := getContent(n)
				price, _ := strconv.Atoi(tmp)
				s.SetPrice(price)
				s.isPriceFind = true
				break
			}
		}
	}
	if !(s.isPriceFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s.isPriceFind == true {
				break
			}
			s.FindPrice(c)
		}
	}
}

// yahoo item info crawler
type yahooItemInfo struct {
	itemInfo
}

func NewYahooItemInfo(n *html.Node) *yahooItemInfo {
	s := new(yahooItemInfo)

	s.isNameFind = false
	s.isUrlFind = false
	s.isPriceFind = false
	s.FindName(n)
	s.FindUrl(n)
	s.FindPrice(n)
	return s
}
func (s *yahooItemInfo) FindName(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "BaseGridItem__title___2HWui" {
				itemName := getContent(n)
				s.isNameFind = true
				s.SetName(itemName)
				break
			}
		}
	}
	if !(s.isNameFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindName(c)
		}
	}
}
func (s *yahooItemInfo) FindUrl(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				s.SetUrl(a.Val)
				s.isUrlFind = true
				break
			}
		}
	}
	if !(s.isUrlFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindUrl(c)
		}
	}
}
func (s *yahooItemInfo) FindPrice(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "BaseGridItem__price___31jkj" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "em" {
						tmp := getContent(c)
						r, _ := regexp.Compile("([0-9]+)")
						match := r.FindString(tmp)
						price, _ := strconv.Atoi(match)
						s.SetPrice(price)
						s.isPriceFind = true
						break
					}
				}
				break
			}
		}
	}
	if !(s.isPriceFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s.isPriceFind == true {
				break
			}
			s.FindPrice(c)
		}
	}
}

// ruten item info crawler
type rutenItemInfo struct {
	itemInfo
}

func NewRutenItemInfo(n *html.Node) *rutenItemInfo {
	s := new(rutenItemInfo)

	s.isNameFind = false
	s.isUrlFind = false
	s.isPriceFind = false
	s.FindName(n)
	s.FindUrl(n)
	s.FindPrice(n)
	return s
}

func (s *rutenItemInfo) FindName(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "h5" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "prod_name " {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "a" {
						name := getContent(c)
						s.SetName(name)
						fmt.Println(name)
						s.isNameFind = true
						break
					}
				}
				break
			}
		}
	}
	if !(s.isNameFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindName(c)
		}
	}
}

func (s *rutenItemInfo) FindUrl(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "h5" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "prod_name " {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "a" {
						for _, a := range c.Attr {
							if a.Key == "href" {
								url := a.Val
								s.SetUrl(url)
								s.isUrlFind = true
								break
							}
						}
						break
					}
				}
				break
			}
		}
	}
	if !(s.isUrlFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindUrl(c)
		}
	}
}

func (s *rutenItemInfo) FindPrice(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "price" {
				tmp := getContent(n)
				r, _ := regexp.Compile("([0-9]+)")
				match := r.FindString(tmp)
				price, _ := strconv.Atoi(match)
				s.SetPrice(price)
				s.isPriceFind = true
				break
			}
		}
	}
	if !(s.isPriceFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s.isPriceFind == true {
				break
			}
			s.FindPrice(c)
		}
	}
}
