package client

import (
	"encoding/json"
	"testing"
)

func TestTrie(t *testing.T) {
	keys := []string{
		"/",
		"/titangrm",
		"/titanscp",
		"/titangrm/test",
		"/titangrm/test/1.1",
		"/titangrm/test/1.2",
		"/titangrm/test/1.1/1.1.1",
	}

	trie := NewTrie("v3", "")
	for _, key := range keys {
		trie.Insert(key)
	}

	content, err := json.MarshalIndent(trie, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("contetn:\n %s", string(content))
}
