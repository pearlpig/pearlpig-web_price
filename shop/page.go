package shop

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type pageParser struct {
	client *http.Client
	target string
}

func NewPageParser(client *http.Client) *pageParser {
	p := new(pageParser)
	p.client = client
	return p
}

func (p *pageParser) SetTarget(target string) {
	p.target = target
}

func (p *pageParser) Parse() *html.Node {
	// connecting the web
	// timeout := time.Duration(5 * time.Second)
	// client := &http.Client{Timeout: timeout}
	// client := &http.Client{}
	req, _ := http.NewRequest("GET", p.target, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := p.client.Do(req)

	if err != nil {
		fmt.Println("Http get err:", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("html parse err", err)
	}
	return doc
}
