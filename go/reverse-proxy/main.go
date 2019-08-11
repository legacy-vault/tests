package main

import "time"

func main() {

	var err error
	var srv Server

	//err = srv.configure("0.0.0.0:2000", "http://localhost:8080")
	err = srv.configure(
		"0.0.0.0:2000",
		"https://yandex.ru/search",
		true,
	)
	if err != nil {
		panic(err)
	}
	srv.start()
	time.Sleep(time.Minute * 5)
	err = srv.stop()
	if err != nil {
		panic(err)
	}
}
