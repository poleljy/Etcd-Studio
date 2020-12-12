package storev3

import (
	"fmt"
	"testing"
	"time"
)

func TestStorev3(t *testing.T) {
	start := time.Now()
	url := "http://139.198.13.149:31339"
	//url := "http://192.168.1.218:30376"

	etcd := DynamicEtcd{
		Endpoints: []string{url},
	}
	if err := etcd.Connect(); err != nil {
		t.Fatal(err)
	}
	defer etcd.DisConnect()

	_, err := etcd.GetKeysOnly("/")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get:", time.Now().Sub(start))

	//trie := client.NewTrie(client.VersionV3, url)
	//for _, key := range keys {
	//	trie.Insert(key)
	//}

	//trieTmp := &trie
	//for _, node := range (*trieTmp).RootNode.Children {
	//	node.Children = nil
	//}

	//content, err := json.MarshalIndent(trie, "", "\t")
	//if err != nil {
	//	t.Fatal(err)
	//}
	////t.Logf("contetn:\n %s", string(content))
	//if ioutil.WriteFile("etcd.json", content, 0644) == nil {
	//	fmt.Println("写入文件成功:", "etcd.json")
	//}
	fmt.Println("all:", time.Now().Sub(start))
}
