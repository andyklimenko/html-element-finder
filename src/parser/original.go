package parser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func doFindOriginal(r io.Reader, rootNode *html.Node, elementID string) (string, *html.Node) {
	for child := rootNode.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode /*&& strings.ToLower(child.Data) == "a" */ {
			for _, a := range child.Attr {
				if strings.ToLower(a.Key) == "id" && a.Val == elementID {
					return child.Data, child
				}
			}
		}

		tagPath, nodeFound := doFindOriginal(r, child, elementID)
		if nodeFound != nil {
			return fmt.Sprintf("%s > %s", child.Data, tagPath), nodeFound
		}
	}

	return "", nil
}

func FindOriginalElement(r io.Reader, elementID string) (string, *html.Node, error) {
	rootNode, err := html.Parse(r)
	if err != nil {
		return "", nil, fmt.Errorf("parsing html: %w", err)
	}

	tagPath, nodeFound := doFindOriginal(r, rootNode, elementID)
	return tagPath, nodeFound, nil
}
