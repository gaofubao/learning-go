package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()

	caCrt, err := ioutil.ReadFile("../cert/ca/ca.pem")
	if err != nil {
		log.Fatal("read ca.crt file error:", err.Error())
	}
	pool.AppendCertsFromPEM(caCrt)

	//cliCrt, err := tls.LoadX509KeyPair("../cert/client/client.pem", "../cert/client/client-key.pem")
	//if err != nil {
	//	log.Fatalln("LoadX509KeyPair error:", err.Error())
	//}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://www.example.net:8080/")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
