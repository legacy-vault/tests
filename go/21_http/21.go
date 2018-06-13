// 21.go.

package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"net/url"
)

func main() {

	var client http.Client
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte
	var srv *http.Server
	var transp *http.Transport

	fmt.Println("HTTP Test.")

	// 1. Simple HTTP 'GET' Request.
	/*
	resp, err = http.Get("https://ya.ru")
	if err != nil {
		log.Fatal(err)
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", resp.Status,
		"Proto:", resp.Proto,
		"ContentLength:", resp.ContentLength,
		"TransferEncoding:", resp.TransferEncoding)

	fmt.Println(string(respBody))
	fmt.Println(respBody)
	*/

	// 2. Advanced HTTP 'GET' Request.
	transp = &http.Transport{}
	transp.DisableCompression = true
	transp.DisableKeepAlives = true
	transp.ExpectContinueTimeout = time.Second * 60
	transp.IdleConnTimeout = time.Second * 60
	client.Timeout = time.Second * 60
	client.Transport = transp
	req, err = http.NewRequest("GET", "https://ya.ru", nil)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0`)

	fmt.Printf("Request: %+v.\r\n", *req)

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %+v.\r\n", *resp)

	fmt.Println("Status:", resp.Status,
		"Proto:", resp.Proto,
		"ContentLength:", resp.ContentLength,
		"TransferEncoding:", resp.TransferEncoding)

	fmt.Println(string(respBody))
	fmt.Println(respBody)

	// 3. Simple HTTP Server.
	srv = &http.Server{}
	srv.Addr = "0.0.0.0:12345"
	srv.Handler = http.HandlerFunc(httpHandler)
	srv.ReadTimeout = time.Second * 60
	srv.WriteTimeout = time.Second * 60

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var reqBody []byte
	var reqURLPath string
	var reqURLParams url.Values

	log.Printf("Request: %+v.\r\n", *r)

	reqBody, err = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	reqURLPath = r.URL.Path
	reqURLParams, err = url.ParseQuery(r.URL.RawQuery)

	log.Println("Request Body (String):", string(reqBody))
	log.Println("Request Body:", reqBody)
	log.Printf("reqURLPath: %+v.\r\n", reqURLPath)
	log.Printf("reqURLParams: %+v.\r\n", reqURLParams)

	// Do something...

	// Send Reply.
	fmt.Fprint(w, "This is a test reply. "+
		"Page '"+ reqURLPath+ "' is not implemented yet. ")
}
