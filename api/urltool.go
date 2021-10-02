package api

import (
	"fmt"
	"net/url"
	"strings"
)

func ResolveRelativeURL(base_url string, href string) string {
	/*
		Resolves Relative URL from href and base url
	*/

	// parse baseurl
	baseurl, _ := url.Parse(base_url)
	base_url = fmt.Sprintf("%s://%s", baseurl.Scheme, baseurl.Host)

	if strings.HasPrefix(href, "/") {
		return fmt.Sprintf("%s%s", base_url, href)
	}
	return href
}
