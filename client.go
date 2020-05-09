package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// 默认超时时间
var timeout = 500 * time.Millisecond
var roundtrip = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
	// DisableKeepAlives: true,
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          300,
	MaxIdleConnsPerHost:   50,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// Client 请求客户端
var Client = &http.Client{Timeout: timeout, Transport: roundtrip}
