package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"

	restful "github.com/emicklei/go-restful/v3"

	"studio/etcd/client"
	rest_api "studio/rest-api"
	"studio/util"
)

func main() {
	port := flag.String("p", "8080", "port")

	flag.CommandLine.Parse(os.Args[1:])

	// restful
	wsContainer := restful.NewContainer()

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders: []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length",
			"Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Allow-Headers"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: true,
		Container:      wsContainer,
	}
	wsContainer.Filter(cors.Filter)

	svc := rest_api.ClientSvc{
		Endpoints: make(map[string]*client.Endpoint),
	}
	wsContainer.Add(svc.WebService())

	// 静态请求，由AssetFS统一处理。
	wsContainer.Handle("/", http.FileServer(AssetFile()))

	// 打开默认浏览器
	name, args := util.GetBrowserCmd(`http://localhost:` + *port)
	exec.Command(name, args...).Start()

	log.Print("start listening on localhost:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, wsContainer))
}
