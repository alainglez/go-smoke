// Smoke crawls web links starting with the smoketest and urls arguments.
//
// This version uses bounded parallelism.
// 
package controllers

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"golang.org/x/net/html"
	"github.com/alainglez/go-smoke/models"
)

//!+
func Smoke(smoketest *models.SmokeTest,  testurls []models.TestUrl) {
	
	// Create visit goroutines to fetch each link.
	for i := 0; i < len(testurls); i++ {
		go func() {
				link := testurls[i].Url
				htmlfragment := testurls[i].HtmlFragment
				hostip := smoketest.Hostip
				statuscode := visit(hostip,link,htmlfragment)
				smoketest.UrlResults[i].Url = link
				smoketest.UrlResults[i].StatusCode = statuscode
				if smoketest.PassFail = "FAIL" {
					continue
				}
				else {
					if statuscode != http.StatusOK {
						smoketest.PassFail = "FAIL"
					}
					else {
					smoketest.PassFail = "PASS"
					}
				}
		}()
	}


// Visit makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func visit(hostip string, url string, htmlfragment string) (string, error) {
	// for now, uri will be url
	uri = url
	resp, err := http.Get(hostip + uri)
	if err != nil {
		return nil, err
	}
	statuscode := resp.StatusCode
	if statuscode != http.StatusOK {
		resp.Body.Close()
		return statuscode, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return statuscode, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	/* No need to return Links
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	*/
	return statuscode, nil
}

//!-Extract

// No need to visit nodes for now
/*
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
*/
//!-
