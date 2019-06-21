package main

import (
	"easy-proxy/admin"
	"easy-proxy/blance"
	"easy-proxy/proxy"
	"log"
	"net/http"
)

func startProxyServer(addr string)  {

	proxy :=  &proxy.ProxyHandle{}
	proxy.Addrs = []string{"127.0.0.1:9999"}
	proxy.Rr = blance.NewWeightedRR(blance.NGINX)

	w := 1
	for _, e := range proxy.Addrs {
		proxy.Rr.Add(e, w)
		w++
	}

	err := http.ListenAndServe(addr, proxy)
	if err == nil {
		log.Panicf("proxy ListenAndServe: %s", addr)
	}

}

func main() {

	go admin.AdminServer(":3000")

	go startProxyServer(":10000")

	select {}

	//go startProxyServer(":10000")
}