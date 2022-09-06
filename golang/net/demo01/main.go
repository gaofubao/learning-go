package main

import (
	"crypto/tls"
	"fmt"
	"math"
	"time"
)

type DomainCheckSSLData struct {
	Domain     string `json:"domain" description:"域名"`
	Record     string `json:"record" description:"记录名"`
	Value      string `json:"value" description:"记录值"`
	Type       string `json:"type" description:"记录类型"`
	Subject    string `json:"subject" description:"证书主题"`
	ExpireDate string `json:"expire_date" description:"过期时间"`
	Remaining  int    `json:"remaining" description:"剩余时间"`
	Error      string `json:"error" description:"错误原因"`
}

func main() {
	data := CheckSSL("", "app", "CNAME", "")
	for _, datum := range data {
		fmt.Println(*datum)
	}
}

func CheckSSL(domain, record, typ, value string) []*DomainCheckSSLData {
	var data []*DomainCheckSSLData

	if record == "*" || record == "@" {
		record = ""
	}
	if typ == "AAAA" {
		value = fmt.Sprintf("[%s]", value)
	}

	addr := fmt.Sprintf("%s:443", value)
	serverName := fmt.Sprintf("%s.%s", record, domain)

	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         serverName,
	})
	if err != nil {
		datum := &DomainCheckSSLData{
			Domain: domain,
			Record: record,
			Value:  value,
			Type:   typ,
			Error:  fmt.Sprintf("检测失败: %v", err),
		}
		data = append(data, datum)
		return data
	}

	state := conn.ConnectionState()
	defer func() { _ = conn.Close() }()
	peerCertificates := state.PeerCertificates
	for _, cert := range peerCertificates {
		notAfter := cert.NotAfter.Local()
		difference := time.Until(notAfter).Seconds()
		datum := &DomainCheckSSLData{
			Domain:     domain,
			Record:     record,
			Value:      value,
			Type:       typ,
			Subject:    cert.Subject.String(),
			ExpireDate: notAfter.Format("2006-01-02 15:04:05"),
			Remaining:  int(math.Floor(difference)),
		}
		data = append(data, datum)
	}

	return data
}
