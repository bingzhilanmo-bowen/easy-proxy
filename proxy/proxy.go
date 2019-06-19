package proxy

import (
	"easy-proxy/blance"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyHandle struct {
	Addrs []string
	Rr blance.RR
}




func (handle *ProxyHandle) ServeHTTP (w http.ResponseWriter, r *http.Request)  {

	addr := handle.Rr.Next().(string)
	remote, err := url.Parse("http://" + addr)
	if err != nil {
		panic(err)
	}
   proxy :=	httputil.NewSingleHostReverseProxy(remote)
   proxy.ServeHTTP(w,r)
}