package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func NewMultipleHostsReverseProxy(target *url.URL) *httputil.ReverseProxy {
	// director := func(req *http.Request) {
	// 	target := targets[rand.Int()%len(targets)]
	// 	req.URL.Scheme = target.Scheme
	// 	req.URL.Host = target.Host
	// 	req.URL.Path = target.Path
	// }

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}

	return &httputil.ReverseProxy{Director: director}
}

func main1() {
	proxy := NewMultipleHostsReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}

var targetAddr *net.TCPAddr

func main2() {
	var target string
	var port int

	flag.StringVar(&target, "remoteSrv", "", "the remote server (<host>:<port>)")
	flag.IntVar(&port, "proxyPort", 7757, "the proxy port")
	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	targetAddr, err = net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Could not start proxy server on %d: %v\n", port, err)
	}

	fmt.Printf("Proxy server running on %d\n", port)
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Could not accept client connection", err)
		}
		go handleTCPConn(conn)
	}
}

func handleTCPConn(conn *net.TCPConn) {
	defer conn.Close()
	log.Printf("Client '%v' connected!\n", conn.RemoteAddr())

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Second * 15)

	client, err := net.DialTCP("tcp", nil, targetAddr)
	if err != nil {
		log.Println("Could not connect to remote server:", err)
		return
	}
	defer client.Close()
	log.Printf("Connection to server '%v' established!\n", client.RemoteAddr())

	client.SetKeepAlive(true)
	client.SetKeepAlivePeriod(time.Second * 15)

	stop := make(chan bool)

	go func() {
		_, err := io.Copy(client, conn)
		log.Println(err)
		stop <- true
	}()

	go func() {
		_, err := io.Copy(conn, client)
		log.Println(err)
		stop <- true
	}()

	<-stop
}
