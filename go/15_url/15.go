// 15.go.

package main

import "fmt"

//import "net"
import "net/url"

func main() {

	var err error
	var urlObj *url.URL
	var urlString string
	var urlValues url.Values

	urlString = "https://user:pass@host.com:5432/path?k=v&id=125&m=15xt&s=%25%24%23#anchor_a"

	urlObj, err = url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("urlObj=%+v\r\n", *urlObj)

	urlValues, err = url.ParseQuery(urlObj.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlValues)
}
