package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))

	container := restful.NewContainer() // 内部包含一个 ServeMux
	container.Add(ws)                   // 内部调用 serveMux.HandleFunc()

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: container}
	log.Fatal(server.ListenAndServe())
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
