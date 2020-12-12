package client

import (
	"fmt"

	"studio/etcd/storev3"
)

const (
	VersionV2 = "v2"
	VersionV3 = "v3"

	TrieMode = "trie"
)

// Etcd Client interface
type Client interface {
	Connect() error
	IsConnect() bool
	DisConnect() error

	// 获取特定前缀的所有key值
	GetKeysOnly(prefix string) ([]string, error)

	// 获取属性值
	GetKeyValue(key string) (string, error)

	// 移除Key
	DeleteKey(key string) error

	// 新增或修改属性值
	PutKeyValue(key, value string, ttl int64) error
}

func NewClient(version, endpoint string) Client {
	var cli Client
	if version == VersionV2 {
		// TODO
	} else if version == VersionV3 {
		cli = &storev3.DynamicEtcd{
			Endpoints: []string{endpoint},
		}
	}
	return cli
}

// Etcd endpoint
type Endpoint struct {
	Url     string
	Version string
	Client  Client
	Trie    *TrieNode
	Keys    []string
}

type Endpoints map[string]*Endpoint

func (end *Endpoint) GetAllKeys(prefix string, trie, refresh bool) (interface{}, error) {
	if !refresh {
		if end.Trie == nil || len(end.Keys) == 0 {
			refresh = true
		}
	}

	// TODO
	refresh = true

	if len(prefix) == 0 {
		prefix = "/"
	}

	if refresh {
		fmt.Println("refresh endpoint:", end.Url, end.Version)
		keys, err := end.Client.GetKeysOnly(prefix)
		if err != nil {
			return nil, err
		}

		trie := NewTrie(end.Version, end.Url)
		for _, key := range keys {
			trie.Insert(key)
		}

		end.Keys = keys
		if prefix == "/" {
			end.Trie = trie.RootNode
		} else if len(trie.RootNode.Children) == 1 {
			end.Trie = trie.RootNode.Children[0]
		}
	}

	if trie {
		return end.Trie, nil
	} else {
		return end.Keys, nil
	}
}
