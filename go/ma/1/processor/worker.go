//============================================================================//
//
// Copyright © 2019 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2019-11-16.
//
//============================================================================//

package processor

// The Worker receives Tasks, performs them and stores the Results.
// Processor may create several Workers depending on its Settings.
// When all the Workers finish their Job, the Processor gathers the Information
// from all the Workers.

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Worker's Errors.
const (
	ErrWorkerIsRunning       = "Worker is already running!"
	ErrWorkerProcessorNotSet = "Worker's Processor is not set!"
)

// Worker is an Object which does the URL Procession Work.
// Several Workers can be used simultaneously.
type Worker struct {
	// Guard against multiple Users.
	lock sync.Mutex

	// Unique Worker's ID. may be used for Logging and Inspection.
	id uint
	// A Link to the Processor.
	processor *Processor
	// The total Number of Pattern Matches that were found by the Worker.
	// This Parameters is set by the Worker and is read by the Processor when
	// the Worker finishes its Job.
	numberOfGoStrings int

	// Internal Flag indicating the Worker's Busy Status.
	isRunning bool
}

// Worker Constructor.
func NewWorker(
	id uint,
	processor *Processor,
) *Worker {
	return &Worker{
		id:                id,
		processor:         processor,
		numberOfGoStrings: 0,
	}
}

// Starts the Worker's main Work Loop.
func (w *Worker) Start() (err error) {

	// Forbid simultaneous Usage.
	w.lock.Lock()
	defer w.lock.Unlock()

	// Fool Check.
	if w.isRunning {
		return errors.New(ErrWorkerIsRunning)
	}
	if w.processor == nil {
		return errors.New(ErrWorkerProcessorNotSet)
	}

	w.isRunning = true
	go w.run()
	return
}

// Worker's main Work Loop.
func (w *Worker) run() {

	var err error
	var httpClient *http.Client
	var task Task

	//log.Printf("Worker %v has started.", w.id) //DEBUG

	// Preparation.
	httpClient = &http.Client{
		Timeout: time.Second * 60,
	}

	// Process Tasks.
	for task = range w.processor.tasksQueue {
		//log.Printf("Worker %v has received a task: %v", w.id, task.UrlAddress) //DEBUG

		// Make an HTTP GET Request and process the Response.
		task.Result.NumberOfGoStrings, err = w.makeAndProcessHttpRequest(
			httpClient,
			task.UrlAddress,
		)
		if err != nil {
			w.processor.workerErrors <- err
			continue
		}

		// Save the Results.
		w.numberOfGoStrings += task.Result.NumberOfGoStrings
	}

	w.isRunning = false
	w.processor.workersWG.Done()
	//log.Printf("Worker %v has stopped.", w.id) //DEBUG
}

// Makes an HTTP GET Request and processes the HTTP Response.
// Returns the Number of 'Go' Strings in the Response using the
// case-insensitive Search.
func (w *Worker) makeAndProcessHttpRequest(
	httpClient *http.Client,
	urlAddress string,
) (numberOfGoStrings int, err error) {

	const (
		patternLowCase = "go"
	)

	var httpResponse *http.Response
	var httpResponseBody []byte

	// Try to make the HTTP GET Request.
	// Do not quit on Error.
	httpResponse, err = httpClient.Get(urlAddress)
	if err != nil {
		return
	}

	// Get the HTTP Response.
	defer func() {
		// Close the HTTP Response Body skipping the Error.
		var derr error
		derr = httpResponse.Body.Close()
		if derr != nil {
			log.Println(derr)
		}
	}()
	httpResponseBody, err = ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	// Count the Pattern String in the HTTP Response Body.
	numberOfGoStrings = w.countPatternStringInHttpBody(
		httpResponseBody,
		patternLowCase,
	)
	fmt.Println(urlAddress, "...", numberOfGoStrings)
	return
}

// Counts the Pattern String in the HTTP Body.
func (w *Worker) countPatternStringInHttpBody(
	httpResponseBody []byte,
	patternLowCase string,
) (numOfPatterns int) {

	var text string

	// Pre-process the input Text.
	text = strings.ToLower(string(httpResponseBody))

	// Counting.
	numOfPatterns = strings.Count(text, patternLowCase)
	return
}

// Returns the total Pattern Matches found by the Worker.
func (w Worker) GetSummary() (numberOfGoStrings int) {
	return w.numberOfGoStrings
}
