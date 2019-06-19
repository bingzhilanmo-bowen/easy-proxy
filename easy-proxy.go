package main

import (
	"easy-proxy/blance"
	"easy-proxy/proxy"
	"log"
	"net/http"
)

func startServer()  {

	proxy :=  &proxy.ProxyHandle{}
	proxy.Addrs = []string{"127.0.0.1:9999"}
	proxy.Rr = blance.NewWeightedRR(blance.NGINX)

	w := 1
	for _, e := range proxy.Addrs {
		proxy.Rr.Add(e, w)
		w++
	}

	err := http.ListenAndServe(":10000", proxy)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}


}


func main() {

	startServer()

}