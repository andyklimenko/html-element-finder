package parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func doMatchWith(r io.Reader, rootNode *html.Node, matchWith *html.Node) (string, float64) {
	matchMap := map[string]float64{}
	for child := rootNode.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			m := nodesMatch(matchWith, child)
			if m > 0 {
				matchMap[child.Data] = m
			}
		}

		tagPath, childMatch := doMatchWith(r, child, matchWith)
		if childMatch > 0 {
			matchMap[fmt.Sprintf("%s > %s", child.Data, tagPath)] = childMatch
		}
	}

	var maxMatch float64
	var resPath string
	for k, v := range matchMap {
		if v > maxMatch {
			maxMatch = v
			resPath = k
		}
	}

	return resPath, maxMatch
}

func FindMatch(r io.Reader, matchWith *html.Node) (string, error) {
	rootNode, err := html.Parse(r)
	if err != nil {
		return "", fmt.Errorf("parsing html: %w", err)
	}

	tagPath, _ := doMatchWith(r, rootNode, matchWith)
	return tagPath, nil
}
