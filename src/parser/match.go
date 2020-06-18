package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func nodesMatch(nodeA *html.Node, nodeB *html.Node) float64 {
	if strings.ToLower(nodeA.Data) != strings.ToLower(nodeB.Data) {
		return 0
	}

	attrsMapA := make(map[string]string)
	for _, a := range nodeA.Attr {
		attrsMapA[a.Key] = a.Val
	}
	attrsMapB := make(map[string]string)
	for _, a := range nodeB.Attr {
		attrsMapB[a.Key] = a.Val
	}

	var matchCount float64
	for k, v := range attrsMapA {
		bAttr, ok := attrsMapB[k]
		if ok && v == bAttr {
			matchCount++
		}
	}

	return matchCount / float64(len(nodeA.Attr))
}
