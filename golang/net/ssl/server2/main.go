package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main() {
	// ssl 双向检验
	pool := x509.NewCertPool()
	crt, err := ioutil.ReadFile("../cert/client/client.pem")
	if err != nil {
		log.Fatalln("读取证书失败！", err.Error())
	}
	pool.AppendCertsFromPEM(crt)

	http.HandleFunc("/", handler)

	s := &http.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert, // 检验客户端证书

		},
	}
	log.Fatal(s.ListenAndServeTLS(
		"../cert/server/server.pem",
		"../cert/server/server-key.pem"),
	)
}
