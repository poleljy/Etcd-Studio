package client

import (
	"fmt"
	"strings"
)

// 树形节点
type TrieNode struct {
	IsParent bool        `json:"isParent"`
	Label    string      `json:"label"`
	Children []*TrieNode `json:"children,omitempty"`

	Path string               `json:"id"`
	next map[string]*TrieNode `json:"next"`
}

type Trie struct {
	RootNode *TrieNode `json:"root"`
}

func NewTrie(version, endpoint string) *Trie {
	label := fmt.Sprintf("%s(%s)", endpoint, version)
	return &Trie{
		RootNode: &TrieNode{
			IsParent: true,
			Label:    label,
			Children: nil,
			Path:     "/",
			next:     make(map[string]*TrieNode),
		},
	}
}

func (tree *Trie) Insert(path string) {
	if len(path) == 0 || path == tree.RootNode.Path {
		return
	}

	cur := tree.RootNode
	keys := strings.Split(path, "/")
	for _, key := range keys {
		if len(key) == 0 || key == tree.RootNode.Path {
			continue
		}

		if _, ok := cur.next[key]; !ok {
			path := cur.Path + "/" + key
			if cur == tree.RootNode {
				path = "/" + key
			}

			node := &TrieNode{
				IsParent: false,
				Label:    key,
				Path:     path,
				next:     make(map[string]*TrieNode),
			}
			cur.IsParent = true
			cur.next[key] = node
			cur.Children = append(cur.Children, node)
		}
		cur = cur.next[key]
	}
}
