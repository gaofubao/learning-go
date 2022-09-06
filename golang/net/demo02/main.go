package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://www.baidu.com:443")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	for _, cert := range resp.TLS.PeerCertificates {
		fmt.Println(cert.Subject, cert.NotAfter.Local().Format("2006-01-02 15:04:05"))
	}
}
