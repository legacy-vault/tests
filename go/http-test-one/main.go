package main

// A simple Network Connection Test.

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/vault-thirteen/random"
)

// Errors.
const (
	ErrfStatusBad = "Bad HTTP Status: %v"
)

// Settings.
const (
	BullshitRequestsCountPerInterval = 200
	BullshitSleepIntervalMs          = 1000
	SearchSleepIntervalS             = 30
)

func main() {
	fmt.Println("Start")
	var errorsReceiver chan error = make(chan error)
	go listenToErrors(errorsReceiver)
	var requests []Request = composeRequests()
	go sendSearchRequests(errorsReceiver, requests)
	go sendBullshitRequests(errorsReceiver)

	// Wait forever.
	var quitChan = make(chan bool)
	<-quitChan
}

func listenToErrors(
	errorsReceiver chan error,
) {
	var err error
	var errorsCount uint
	for err = range errorsReceiver {
		errorsCount++
		if errorsCount%1000 == 0 {
			log.Println(time.Now().Format(time.Stamp), "> errorsCount:", errorsCount)
		}

		// Error Filtering.
		if isErrorOfUnreachableHost(err) {
			continue
		} else {
			log.Println(time.Now().Format(time.Stamp), ">", err)
		}
	}
}

func isErrorOfUnreachableHost(
	err error,
) (result bool) {
	const (
		errPatternA = "A socket operation was attempted to an unreachable network."
		errPatternB = "The requested address is not valid in its context."
		errPatternC = "A socket operation was attempted to an unreachable network."
		errPatternD = "No connection could be made because the target machine actively refused it."
		errPatternE = " A connection attempt failed because the connected party did not properly respond" +
			" after a period of time, or established connection failed because connected host has failed to respond."
	)
	var errText string = err.Error()
	if strings.Contains(errText, errPatternA) {
		result = true
		return
	}
	if strings.Contains(errText, errPatternB) {
		result = true
		return
	}
	if strings.Contains(errText, errPatternC) {
		result = true
		return
	}
	if strings.Contains(errText, errPatternD) {
		result = true
		return
	}
	if strings.Contains(errText, errPatternE) {
		result = true
		return
	}
	result = false
	return
}

func sendSearchRequests(
	errorsReceiver chan error,
	requests []Request,
) {
	// Cache the full URLs.
	var fullUrls []string = make([]string, 0)
	for _, request := range requests {
		fullUrls = append(
			fullUrls,
			request.BaseUrl+"?"+request.SearchQueryText,
		)
	}

	// Periodically send the Requests.
	var sleepInterval time.Duration = time.Second * SearchSleepIntervalS
	for {
		// Send a Series of Search Requests.
		var uniqueRequestsCount = len(requests)
		for i := 0; i < uniqueRequestsCount; i++ {
			go sendSearchRequest(errorsReceiver, &fullUrls[i])
		}

		// Wait.
		time.Sleep(sleepInterval)
	}
}

func sendSearchRequest(
	errorsReceiver chan error,
	url *string, // Must be non-null.
) {
	var httpClient *http.Client = new(http.Client)
	var err error
	var resp *http.Response
	resp, err = httpClient.Get(*url)
	if err != nil {
		errorsReceiver <- err
		return
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf(ErrfStatusBad, resp.Status)
			errorsReceiver <- err
			return
		}
	}
}

func sendBullshitRequests(
	errorsReceiver chan error,
) {
	// Periodically send the Requests.
	var sleepInterval time.Duration = time.Millisecond * BullshitSleepIntervalMs
	for {
		// Send a Series of Bullshit Requests.
		for i := 0; i < BullshitRequestsCountPerInterval; i++ {
			go sendBullshitRequest(errorsReceiver)
		}

		// Wait.
		time.Sleep(sleepInterval)
	}
}

func sendBullshitRequest(
	errorsReceiver chan error,
) {
	var err error
	var url string
	url, err = makeRandomHttpUrl()
	if err != nil {
		errorsReceiver <- err
		return
	}

	var httpClient *http.Client = new(http.Client)
	var resp *http.Response
	resp, err = httpClient.Get(url)
	if err != nil {
		errorsReceiver <- err
		return
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf(ErrfStatusBad, resp.Status)
			errorsReceiver <- err
			return
		}
	}
}

func makeRandomHttpUrl() (result string, err error) {
	var s1, s2, s3, s4 uint
	s1, err = random.Uint(0, 255)
	if err != nil {
		return
	}
	s2, err = random.Uint(0, 255)
	if err != nil {
		return
	}
	s3, err = random.Uint(0, 255)
	if err != nil {
		return
	}
	s4, err = random.Uint(0, 255)
	if err != nil {
		return
	}

	result = "http://" +
		strconv.FormatUint(uint64(s1), 10) + "." +
		strconv.FormatUint(uint64(s2), 10) + "." +
		strconv.FormatUint(uint64(s3), 10) + "." +
		strconv.FormatUint(uint64(s4), 10)
	return
}
