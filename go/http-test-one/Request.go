package main

import "net/url"

type Request struct {
	BaseUrl         string
	SearchQueryText string
}

func composeRequests() (requests []Request) {

	const (
		KeyText = "text"
		KeyQ    = "q"
		//
		ValueTextA = "Шoк! Ceнcaция! Paccлeдoвaниe PEH-TB! Bнимaниe! Achtung! " +
			"Инoплaнeтянe пoxитили Си Цзиньпинa и Влaдимиpa Пyтинa!"
	)

	requests = make([]Request, 0, 2)

	// Search Query Text for Yandex.
	{
		var urlValuesForYandex url.Values = make(url.Values)
		urlValuesForYandex.Add(KeyText, ValueTextA)
		requests = append(requests, Request{
			BaseUrl:         "https://yandex.ru/search",
			SearchQueryText: urlValuesForYandex.Encode(),
		})
	}

	// Search Query Text for Google.
	{
		var urlValuesForGoogle url.Values = make(url.Values)
		urlValuesForGoogle.Add(KeyQ, ValueTextA)
		requests = append(requests, Request{
			BaseUrl:         "https://www.google.com/search",
			SearchQueryText: urlValuesForGoogle.Encode(),
		})
	}

	// Search Query Text for Bing.
	{
		var urlValuesForBing url.Values = make(url.Values)
		urlValuesForBing.Add(KeyQ, ValueTextA)
		requests = append(requests, Request{
			BaseUrl:         "https://www.bing.com/search",
			SearchQueryText: urlValuesForBing.Encode(),
		})
	}

	return
}
