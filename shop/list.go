package shop

import (
	"fmt"
	"regexp"

	"golang.org/x/net/html"
)

type itemList struct {
	rootNode   *html.Node
	listNode   *html.Node
	itemNode   []*html.Node
	isListFind bool
	isItemFind bool
}

func NewItemList() *itemList {
	s := new(itemList)
	s.isListFind = false
	s.isItemFind = false
	return s
}

func (s *itemList) SetRoot(n *html.Node) {
	s.rootNode = n
}

func (s *itemList) SetItem(n *html.Node) {
	s.itemNode = append(s.itemNode, n)
}

func (s *itemList) GetItem() []*html.Node {
	return s.itemNode
}

func (s *itemList) SetList(n *html.Node) {
	s.listNode = n
}

func (s *itemList) GetList() *html.Node {
	return s.listNode
}

type listCrawler interface {
	GetItem() []*html.Node
}

// Add new list crawler method below.
//
func NewListCrawler(s string, n *html.Node) listCrawler {
	switch s {
	case "shopee":
		fmt.Println("create shopee list crawler")
		return NewShopeeItemList(n)
	case "yahoo":
		fmt.Println("create yahoo list crawler")
		return NewYahooItemList(n)
	case "ruten":
		fmt.Println("create ruten list crawler")
		return NewRutenItemList(n)
	// Add new list crawler constructor.
	default:
		panic("unknown server name")
	}
}

// shopee item list crawler
type shopeeItemList struct {
	itemList
}

func NewShopeeItemList(n *html.Node) *shopeeItemList {
	s := new(shopeeItemList)
	s.isListFind = false
	s.SetRoot(n)
	s.FindList(s.rootNode)
	s.FindItem(s.listNode)

	return s
}

func (s *shopeeItemList) FindList(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "row shopee-search-item-result__items" {
				s.isListFind = true
				s.SetList(n)
				break
			}
		}
	}
	if !(s.isListFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindList(c)
		}
	}
}

func (s *shopeeItemList) FindItem(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "col-xs-2-4 shopee-search-item-result__item" {
				s.isItemFind = true
				s.SetItem(n)
				break
			}
		}
	}
	if !(s.isItemFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindItem(c)
		}
	}
}

// yahoo item list crawler
type yahooItemList struct {
	itemList
}

func NewYahooItemList(n *html.Node) *yahooItemList {
	s := new(yahooItemList)
	s.isListFind = false
	s.SetRoot(n)
	s.FindList(s.rootNode)
	s.FindItem(s.listNode)
	return s
}

func (s *yahooItemList) FindList(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "ResultList_mainItemList_29xnt" {
				s.isListFind = true
				s.SetList(n)
				break
			}
		}
	}
	if !(s.isListFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindList(c)
		}
	}
}

func (s *yahooItemList) FindItem(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "li" {
		for _, a := range n.Attr {
			if a.Key == "data-imprsn" {
				r, _ := regexp.Compile("sr")
				if r.MatchString(a.Val) == true {
					s.isItemFind = true
					s.SetItem(n)
					break
				}

			}
		}
	}
	if !(s.isItemFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindItem(c)
		}
	}
}

// ruten item list crawler
type rutenItemList struct {
	itemList
}

func NewRutenItemList(n *html.Node) *rutenItemList {
	s := new(rutenItemList)
	s.isListFind = false
	s.SetRoot(n)
	s.FindList(s.rootNode)
	s.FindItem(s.listNode)
	return s
}

func (s *rutenItemList) FindList(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "dl" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "search_form s_grid" {
				s.isListFind = true
				s.SetList(n)
				break
			}
		}
	}
	if !(s.isListFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindList(c)
		}
	}
}

func (s *rutenItemList) FindItem(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "dl" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "site_loop" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "dd" {
						for _, a := range c.Attr {
							if a.Key == "_item" && a.Val == "nor" {
								s.isItemFind = true
								s.SetItem(c)
								break
							}
						}
					}
				}
			}
		}
	}
	if !(s.isItemFind == true) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			s.FindItem(c)
		}
	}
}
