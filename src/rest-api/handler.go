package rest_api

import (
	restful "github.com/emicklei/go-restful/v3"
)

func (s ClientSvc) getAllKeys(req *restful.Request, res *restful.Response) {
	endpoint, err := s.getEndpoint(req)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	refresh := req.QueryParameter("refresh")
	prefix := req.QueryParameter("prefix")
	trie := req.QueryParameter("trie")
	ret, err := endpoint.GetAllKeys(prefix, trie == "true", refresh == "true")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, ret)
}

func (s ClientSvc) getValue(req *restful.Request, res *restful.Response) {
	key := req.QueryParameter("key")

	endpoint, err := s.getEndpoint(req)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	value, err := endpoint.Client.GetKeyValue(key)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	ret := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   key,
		Value: value,
	}
	ResWriteEntity(res, ret)
}

func (s ClientSvc) addKey(req *restful.Request, res *restful.Response) {
	endpoint, err := s.getEndpoint(req)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	body := addKey{}
	if err := req.ReadEntity(&body); err != nil {
		ResWriteError(res, err)
		return
	}
	if err := endpoint.Client.PutKeyValue(body.Key, body.Value, body.TTL); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, body)
}

func (s ClientSvc) updateValue(req *restful.Request, res *restful.Response) {
	endpoint, err := s.getEndpoint(req)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	body := updateValue{}
	if err := req.ReadEntity(&body); err != nil {
		ResWriteError(res, err)
		return
	}
	if err := endpoint.Client.PutKeyValue(body.Key, body.Value, -1); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, body)
}

func (s ClientSvc) delKey(req *restful.Request, res *restful.Response) {
	key := req.QueryParameter("key")

	endpoint, err := s.getEndpoint(req)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if err := endpoint.Client.DeleteKey(key); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteNullEntity(res)
}
