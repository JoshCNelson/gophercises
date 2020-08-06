package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

type Links []Link

func (l *Link) extractHref(n *html.Node) {
	if len(n.Attr) > 0 {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				l.Href = attr.Val
			}
		}
	}
}

func (l *Link) extractText(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if c.Type == 1 {
			l.Text = strings.TrimSpace(l.Text + c.Data)
		}

		l.extractText(c)
	}
}

func main() {
	file, err := ioutil.ReadFile("./ex1.html")

	if err != nil {
		fmt.Println("Could not find file")
		return
	}

	r := bytes.NewReader(file)

	doc, _ := html.Parse(r)

	anchors := Links{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			anchor := Link{}
			anchor.extractHref(n)
			anchor.extractText(n)

			anchors = append(anchors, anchor)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for i, a := range anchors {
		fmt.Printf("Anchor %d:\nhref: %s\ntext: %s\n", i+1, a.Href, a.Text)
	}
}
