package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"sort"
)

type ReqInfo struct {
	Item     string   `json:"item,omitempty"`
	Server   []string `json:"server,omitempty"`
	NeedSort bool     `json:"need_sort,omitempty"`
}

type ItemInfos struct {
	Server string `json:"server,omitempty"`
	Name   string `json:"name,omitempty"`
	Url    string `json:"url,omitempty"`
	// empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero.
	// use pointer to deal with 0 values
	Price *int `json:"price,omitempty"`
}

type Crawler struct {
	ReqServer []Server
	ReqItem   string
	NeedSort  bool
	PagesNode []Page
	Items     []Item
}

func App(r string) []byte {
	fmt.Println(r)
	req := new(ReqInfo)
	json.Unmarshal([]byte(r), &req)
	c := NewCrawler(req)
	c.Crawl()
	// for _, v := range c.Items {
	// 	fmt.Println("---------")
	// 	fmt.Println("name:", v.Info.Name)
	// 	fmt.Println("price:", v.Info.Price)
	// 	fmt.Println("url:", v.Info.Url)
	// }
	// Sort the item by price
	sort.Sort(c)
	list, _ := json.Marshal(c.Items)

	return list

}
func NewCrawler(req *ReqInfo) *Crawler {
	c := new(Crawler)
	c.GetServerInfo(req.Server)
	c.ReqItem = req.Item
	c.NeedSort = req.NeedSort
	c.initPagesNode()
	return c
}
func (c *Crawler) initPagesNode() {
	nodeNum := 0
	for _, i := range c.ReqServer {
		nodeNum += i.PageNum
	}
	c.PagesNode = make([]Page, 0, nodeNum)
}

func (c *Crawler) Crawl() {
	c.CrawlPage()
	c.CrawlList()
	c.CrawlItem()
}

func (c *Crawler) GetServerInfo(reqServer []string) {
	// get all server info
	serverNum := len(reqServer)
	c.ReqServer = make([]Server, 0, serverNum)
	s := ServerInfo()

	for _, i := range s.List {
		for _, s := range reqServer {
			if s == i.Name {
				c.ReqServer = append(c.ReqServer, i)
			}
		}
	}
	if !(len(c.ReqServer) == serverNum) {
		panic(fmt.Sprintf("Wrong request server number!\nRequest server: %v\nHandle server:%v", reqServer, c.ReqServer))
	}
}

type Info struct {
	Home   string `json:"home,omitempty"`
	Search string `json:"search,omitempty"`
	Page   string `json:"page,omitempty"`
	Sort   string `json:"sort,omitempty"`
}

type Server struct {
	Name    string `json:"name,omitempty"`
	Url     Info   `json:"url,omitempty"`
	PageNum int    `json:"page_num,omitempty"`
}
type AllServerInfo struct {
	List []Server `json:"server,omitempty"`
}

func ServerInfo() *AllServerInfo {
	const dataFile = "server.json"
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), dataFile)
	fmt.Println("path:", datapath)
	s := new(AllServerInfo)
	// jsonFile, err := os.Open("D:/workspace/Golang/src/pearlpig/web_price/dev/one/server.json")
	jsonFile, err := os.Open(datapath)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &s)
	if len(s.List) == 0 {
		panic("Empty server information! Didn't read the json file!")
	}

	return s
}
func Req() *ReqInfo {
	r := new(ReqInfo)
	jsonFile, err := os.Open("D:/workspace/Golang/src/pearlpig/web_price/dev/one/req.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &r)
	return r
}

// Sort package interface init
//Len()
func (c *Crawler) Len() int {
	return len(c.Items)
}

//Less():
func (c *Crawler) Less(i, j int) bool {
	return *c.Items[i].Info.Price < *c.Items[j].Info.Price
}

//Swap()
func (c *Crawler) Swap(i, j int) {
	c.Items[i], c.Items[j] = c.Items[j], c.Items[i]
}
