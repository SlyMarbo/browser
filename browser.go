package browser

import (
	"regexp"
	"strings"

	"exp/html"
)

func css(s string) []string {
	r := regexp.MustCompile(`(?:url\(\s*[\'\"]?)(.*?)(?:[\'\"]?\s*\))`)
	matches := r.FindAllStringSubmatch(s, -1)
	out := make([]string, 0, len(matches))
	for _, match := range matches {
		out = append(out, match[1])
	}
	return out
}

func Links(s string) ([]string, error) {
	out := make([]string, 0, 5)
	
	n, err := html.Parse(strings.NewReader(s))
	if err != nil {
		out = css(s)
		return nil, err
	}
	
	var match func(*html.Node, string, string)
	match = func(n *html.Node, node, field string) {
		if n.Type == html.ElementNode && n.Data == node {
			for _, a := range n.Attr {
				if a.Key == field {
					out = append(out, a.Val)
					return
				}
			}
		}
	}
	
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			match(n, "link", "href")
			match(n, "script", "src")
			if n.Data == "style" {
				out = append(out, css(n.FirstChild.Data)...)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	
	if len(out) == 0 {
		out = css(s)
	}
	
	return out, nil
}
