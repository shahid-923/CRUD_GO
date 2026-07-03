package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

/*
Server is an INTERFACE (contract).
Any type that implements these 3 methods is a Server.
*/
type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

/*
simpleServer is a CONCRETE type.
It becomes a Server automatically because it implements all methods.
*/
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy // forwards requests to backend server
}

/*

returns as Server interface
*/
func newSimpleServer(addr string) Server {
	serverUrl, err := url.Parse(addr) // string → structured URL
	handleError(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

/*
LoadBalancer holds:
- list of servers
- round robin counter
- mutex for concurrency safety
*/
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
	mu              sync.Mutex // protects roundRobinCount
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:    port,
		servers: servers,
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

/* ---- simpleServer implements Server interface ---- */

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true // placeholder (no real health check yet)
}

/*
Serve:
- receives request from load balancer
- ReverseProxy forwards it to real backend
- response comes back to client automatically
*/
func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

/*
Round-robin logic:
- mutex prevents race when many requests come together
- % ensures index cycles back to 0
*/
func (lb *LoadBalancer) getNextAvailableServer() Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	lb.roundRobinCount++

	return server
}

/*
ServeProxy:
- entry point for ALL incoming HTTP requests
- chooses server
- forwards request
*/
func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}

func main() {
	// backend servers
	servers := []Server{
		newSimpleServer("https://www.duckduckgo.com"),
		newSimpleServer("https://bing.com"),
		newSimpleServer("https://search.yahoo.com"),
	}

	// create load balancer
	lb := NewLoadBalancer("9000", servers) 

	// every request goes to ServeProxy
	http.HandleFunc("/", lb.ServeProxy)

	fmt.Printf("Load Balancer running at http://localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
