package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	httpServer           http.Server
	proxyServerTargetURL *url.URL
	proxyServer          *httputil.ReverseProxy
	allowSSL             bool //TODO
}

func (s *Server) configure(
	listenAddress string,
	targetAddress string,
	allowSSL bool,
) error {

	var err error

	s.httpServer = http.Server{
		Addr:    listenAddress,
		Handler: http.HandlerFunc(s.httpRouter),
	}
	s.proxyServerTargetURL, err = url.Parse(targetAddress)
	if err != nil {
		return err
	}

	s.proxyServer = httputil.NewSingleHostReverseProxy(s.proxyServerTargetURL)
	s.proxyServer.ModifyResponse = s.httpProxyPostProcessor
	s.proxyServer.ErrorHandler = s.httpProxyErrorHandler

	s.allowSSL = allowSSL

	return nil
}

func (s *Server) start() {
	go func() {
		var err error
		err = s.httpServer.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()
}

func (s *Server) stop() error {
	return s.httpServer.Shutdown(context.Background())
}

func (s *Server) httpRouter(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Scheme, r.URL.Path) //!

	// Update the Headers to allow for SSL Redirection.
	if s.allowSSL { //TODO
		r.URL.Host = s.proxyServerTargetURL.Host
		r.URL.Scheme = s.proxyServerTargetURL.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = s.proxyServerTargetURL.Host
	}

	s.proxyServer.ServeHTTP(w, r)
}

func (s *Server) httpHandlerError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (s *Server) httpProxyPostProcessor(resp *http.Response) error {

	resp.Header.Add("X-Abc", "Processed by me")
	return nil
}

func (s *Server) httpProxyErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("Proxy Error:", err)
	w.WriteHeader(http.StatusBadGateway)
}
