package rest_api

import (
	api "github.com/emicklei/go-restful-openapi"
	restful "github.com/emicklei/go-restful/v3"

	"studio/etcd/client"
)

type ClientSvc struct {
	Endpoints client.Endpoints
}

// WebService creates a new service that can handle REST requests for resources.
func (s ClientSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api").Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{"etcd api"}

	// 获取所有key
	ws.Route(ws.GET("/{version}/keys").To(s.getAllKeys).
		Doc("get all keys").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("endpoint", "etcd endpoint").DataType("string").Required(true)).
		Param(ws.QueryParameter("refresh", "refresh").DataType("bool").Required(false)).
		Param(ws.QueryParameter("prefix", "prefix").DataType("string").Required(true)).
		Param(ws.QueryParameter("trie", "trie").DataType("bool").Required(false)))

	// 获取key-value
	ws.Route(ws.GET("/{version}/value").To(s.getValue).
		Doc("get value of key").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("endpoint", "etcd endpoint").DataType("string").Required(true)).
		Param(ws.QueryParameter("key", "key").DataType("string").Required(true)))

	// 编辑value
	ws.Route(ws.PUT("/{version}/value").To(s.updateValue).
		Doc("update value of key").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("endpoint", "etcd endpoint").DataType("string").Required(true)).
		Reads(updateValue{}))

	// 添加key
	ws.Route(ws.POST("/{version}/key").To(s.addKey).
		Doc("add key").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("endpoint", "etcd endpoint").DataType("string").Required(true)).
		Reads(addKey{}))

	// 移除key
	ws.Route(ws.DELETE("/{version}/key").To(s.delKey).
		Doc("delete key").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("endpoint", "etcd endpoint").DataType("string").Required(true)).
		Param(ws.QueryParameter("key", "key").DataType("string").Required(true)))

	return ws
}
