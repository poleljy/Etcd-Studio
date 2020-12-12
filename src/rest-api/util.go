package rest_api

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"

	"studio/etcd/client"
)

type addKey struct {
	Key   string `json:'key'`
	Value string `json:'value'`
	TTL   int64  `json:'ttl'`
}

type updateValue struct {
	Key   string `json:'key'`
	Value string `json:'value'`
}

func (s ClientSvc) getEndpoint(req *restful.Request) (*client.Endpoint, error) {
	version := req.PathParameter("version")
	url := req.QueryParameter("endpoint")
	key := fmt.Sprintf("%s_%s", url, version)

	if _, ok := s.Endpoints[key]; !ok {
		cli := client.NewClient(version, url)
		if err := cli.Connect(); err != nil {
			return nil, err
		}

		end := client.Endpoint{
			Url:     url,
			Version: version,
			Client:  cli,
		}
		s.Endpoints[key] = &end
	}
	return s.Endpoints[key], nil
}

////////////////////////////////////////////////////////////////
func ResWriteError(res *restful.Response, err error) error {
	log.Print(err)
	res.Header().Add("Content-Type", "application/json")
	return res.WriteError(http.StatusOK, NewError(err.Error(), http.StatusInternalServerError))
}

func ResWriteEntity(res *restful.Response, value interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	//if isNil(value) {
	//	return res.WriteEntity(nullRet{})
	//}
	return res.WriteEntity(value)
}

type nullRet struct {
	ret string `json:"ret,omitempty"`
}

func ResWriteNullEntity(res *restful.Response) error {
	res.Header().Add("Content-Type", "application/json")
	return res.WriteEntity(nullRet{})
}
