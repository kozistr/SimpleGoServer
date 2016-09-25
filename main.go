/*******************************************
* This program is for testing HTTP2 protocol
*
*
* Project : Simple HTTP2 server
* version : 0.0.1
* lastest : 2016/9/24
* Author  : kozistr
*********************************************/
package main

import (
	"fmt"
	"golang.org/x/net/http2"

	/* Notice that
	Users who use GO version under 1.6
	can't use HTTP2 directly from net/http
	Link : https://godoc.org/golang.org/x/net/http2
	*/
	"net/http" // For HTTP2
)

// SSL files are needed for HTTP2 connection
var (
	cert = "kozistr.cert"
	key  = "kozistr.key"
)

type Context struct {
	// parameters
	paras map[string]interface{}

	resw http.ResponseWriter
	req  *http.Request
}

func Server_Init() http.Server {
	var s http.Server

	// Port 8080
	s.Addr = ":8080"

	// Enable HTTP2
	http2.ConfigureServer(&s, nil)

	return s
}

func main() {
	// Init Server Setting
	s := Server_Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ShowHeader(w, r)
	})

	/* Notice that
	It'll not work without ssl cert.
	So u just need 2 files which .cert and .key
	*/
	s.ListenAndServeTLS(cert, key)
}

func ShoHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Transfer Encoding : %s\n", r.TransferEncoding)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)
}
