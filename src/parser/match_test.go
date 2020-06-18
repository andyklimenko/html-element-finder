package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestNodesSimilar(t *testing.T) {
	t.Parallel()

	originalNode := &html.Node{
		Data: "a",
		Attr: []html.Attribute{
			{Key: "id", Val: "make-everything-ok-button"},
			{Key: "class", Val: "btn btn-success"},
			{Key: "title", Val: "Make-Button"},
			{Key: "href", Val: "#ok"},
		},
	}

	t.Run("different tags", func(t *testing.T) {
		assert.Zero(t, nodesMatch(originalNode, &html.Node{Data: "b"}))
		assert.Zero(t, nodesMatch(originalNode, &html.Node{Data: "A"}))
	})
	t.Run("similar 1", func(t *testing.T) {
		testNode := &html.Node{
			Attr: []html.Attribute{
				{Key: "class", Val: "btn"},
				{Key: "title", Val: "Make-Button"},
				{Key: "href", Val: "#ok"},
			},
			Data: "a",
		}

		assert.InDelta(t, 1, nodesMatch(originalNode, testNode), 0.5)
	})
}
