# 获取key值列表前缀树
## URL: GET v3/keys?endpoint=xxxx&trie=true&refresh=true

##### 参数： 
* [必填] endpoint : etcd服务地址
* [选填] trie : 前缀树模式
* [选填] refresh : 是否强制刷新

# 获取key值对应value
## URL: GET v3/value?endpoint=xxxx&key=xxxx

##### 参数： 
* [必填] endpoint : etcd服务地址
* [必填] key : key值

# 更新value
## URL: PUT v3/value?endpoint=xxxx

##### 参数： 
* [必填] endpoint : etcd服务地址
* [必填] body: 
{
	"key":"xxxx",
	"value":"xxxx"
}

# 添加key-value
## URL: POST v3/key?endpoint=xxxx

##### 参数： 
* [必填] endpoint : etcd服务地址
* [必填] body: 
{
	"key":"xxxx",
	"value":"xxxx"
	"ttl": xx          // ttl：整数，单位是秒，不设置ttl传值 -1
}

# 移除key
## URL: DELETE v3/key?endpoint=xxxx&key=xxxx

##### 参数： 
* [必填] endpoint : etcd服务地址
* [必填] key : key值

