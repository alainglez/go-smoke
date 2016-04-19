// Smoke crawls web links starting with the smoketest and urls arguments.
//
// This version uses bounded parallelism.
// 
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	
	"github.com/alainglez/go-smoke/models"
)

//!+
func Smoke(smoketest *models.SmokeTest,  testurls []models.TestUrl) string {
	// Create visit goroutines to fetch each link.
	for i := 0; i < len(testurls)-1; i++ {
		go func() {
				statusCode, err := visit(&testurls[i])
				smoketest.UrlResults[i].Url = testurls[i].Url
				smoketest.UrlResults[i].StatusCode = statusCode
				if smoketest.PassFail == "FAIL" {
					continue
				} else {
					if statusCode != http.StatusOK {
						smoketest.PassFail = "FAIL"
					} else {
					smoketest.PassFail = "PASS"
					}
				}
		}()
	}
	return smoketest.PassFail
}

// visit makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func visit(testurl *models.TestUrl) (int, error) {
	//hostip string, decodedurl string, htmlfragment string
	var statuscode int
	// replace domain with ip :80 | 443 depending if url has https or not
	u, err := url.Parse(testurl.Url)
	if err != nil {
		return statuscode, err
	}
	if u.Scheme == "http" {
		u.Host = testurl.HostIp + ":80" 
	}
	if u.Scheme == "https" {
		u.Host = testurl.HostIp + ":443"
	}	
	resp, err := http.Get(u)
	if err != nil {
		return statuscode, err
	}
	statuscode = resp.StatusCode
	defer resp.Body.Close()
	if statuscode != http.StatusOK {
		return statuscode, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	if strings.Contains(resp.Body, htmlfragment) {
		return statuscode, nil
	}
	return statuscode, fmt.Errorf("couldn't find %s in %s", htmlfragment, url)
}	
//!-
