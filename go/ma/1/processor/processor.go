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

// The Processor reads URLs from the standard Input Stream and processes them.
// URLs must be separated by ASCII LF Character. An EOF Symbol must be put to
// the End of the Stream to get the processed Data. When the Processing is
// finished, the resulting Data is written to the standard Output Stream
// (StdOut). Errors are written to the standard Error Stream (StdErr).
//
// Notes.
// To simulate the EOF Character from Console in Windows and Ubuntu Operating
// Systems, 'Ctrl+D' Keys Combination can be used.

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

// Processor's Errors & Messages.
const (
	ErrMsgReadingStdin    = "Encountered an Error while reading the standard Input Stream:"
	MsgAborting           = "Aborting the Reading..."
	MsgPatternsTotalCount = "Patterns Total Count:"
)

// Processor's Settings.
const (
	TasksQueueSizeDefault          = 100
	WorkerErrorsChannelSizeDefault = 100
	WorkersCountLimitDefault       = 5
	ProtocolDefault                = "http"
	ProtocolSeparator              = "://"
	LineDelimiter                  = '\n'
)

// Worker is the main Object which controls the URL Reading and Processing.
type Processor struct {
	// Guard against multiple Users.
	lock sync.Mutex

	// Queue of Tasks.
	tasksQueue chan Task
	// Received Tasks Counter.
	// Is used only in Economy Mode.
	receivedTasksCount uint

	// The maximum Number of active Workers allowed.
	workersCountLimit int
	// Economy Mode Indicator.
	// When the Number of Tasks is less then the maximum allowed Workers Count,
	// we are in the Economy Mode. As soon as we reach the maximum allowed
	// Workers Count, we disable the Economy Mode. This Flag is used to
	// analyze whether we need to start additional Workers or not.
	economyModeIsUsed bool
	// List of active Workers.
	workers []*Worker
	// Errors which Workers encounter.
	workerErrors chan error
	// Workers Wait Group.
	workersWG sync.WaitGroup

	// Various cached Settings...
	// Default Protocol Prefix.
	protocolPrefixDefault string
}

// Main public Method of the Processor which does all the Work.
func (p *Processor) Use() {

	var err error
	var numberOfGoStrings int

	// Forbid simultaneous Usage.
	p.lock.Lock()
	defer p.lock.Unlock()

	// Initialization.
	p.init()

	// Process the Stream.
	numberOfGoStrings, err = p.processLinesFromStdin()
	if err != nil {
		log.Println(ErrMsgReadingStdin, err)
		log.Println(MsgAborting)
	}
	fmt.Println(MsgPatternsTotalCount, numberOfGoStrings)
}

// Processor Initialization.
func (p *Processor) init() {
	p.tasksQueue = make(chan Task, TasksQueueSizeDefault)
	p.workersCountLimit = WorkersCountLimitDefault
	p.economyModeIsUsed = true

	// Errors Channel & Listener.
	p.workerErrors = make(chan error, WorkerErrorsChannelSizeDefault)
	go p.errorListener()

}

// Error Listener.
func (p *Processor) errorListener() {

	var err error

	for err = range p.workerErrors {
		log.Println(err)
	}
}

// Reads Lines from the standard Input Stream and processes them.
// Lines must be separated by the ASCII LF Character ('\n').
func (p *Processor) processLinesFromStdin() (totalNumberOfGoStrings int, err error) {

	var mustStop bool
	var readerOfStdin *bufio.Reader
	var urlAddress string
	var worker *Worker

	// Preparation.
	p.protocolPrefixDefault = ProtocolDefault + ProtocolSeparator
	readerOfStdin = bufio.NewReader(os.Stdin)

	// Read the Lines and process them.
	// An EOF triggers the Loop's Stop, an Error triggers a Shutdown.
	for {
		urlAddress, mustStop, err = p.getUrlFromStdin(readerOfStdin)
		if err != nil {
			return
		}
		if mustStop {
			break
		}

		p.sendNewTask(urlAddress)

		err = p.startWorkerIfNeeded()
		if err != nil {
			return
		}
	}

	// Wait for all the Workers to finish their Work. Get the Summary.
	close(p.tasksQueue)
	p.workersWG.Wait()
	for _, worker = range p.workers {
		totalNumberOfGoStrings += worker.GetSummary()
	}

	return
}

// Gets an URL Address from the Reader.
func (p Processor) getUrlFromStdin(
	readerOfStdin *bufio.Reader,
) (urlAddress string, mustStop bool, err error) {

	var line string
	var urlFromLine *url.URL

	// Get a next Line separated by the ASCII 'LF' Character.
	// This Delimiter may be useless for some exotic Operating Systems,
	// e.g. for ancient Commodore, ZX Spectrum and many others.
	// See 'https://en.wikipedia.org/wiki/Newline' for more Information.
	line, err = readerOfStdin.ReadString(LineDelimiter)
	if err != nil {
		if err == io.EOF {
			err = nil
			mustStop = true
		}
		return
	}
	line = strings.TrimSpace(line)

	// Try to convert the Text Line to an URL.
	urlFromLine, err = url.Parse(line)
	if err != nil {
		return
	}

	// Compose the full URL Address.
	// Use the default Protocol if it is not specified.
	if len(urlFromLine.Scheme) > 0 {
		urlAddress = urlFromLine.String()
	} else {
		urlAddress = p.protocolPrefixDefault + urlFromLine.String()
	}
	return
}

// Creates a new Task and puts it into a Queue.
func (p *Processor) sendNewTask(
	urlAddress string,
) {
	var task Task

	task = Task{
		UrlAddress: urlAddress,
	}
	p.tasksQueue <- task
}

// Start a new Worker if necessary.
func (p *Processor) startWorkerIfNeeded() (err error) {

	var worker *Worker

	if !p.economyModeIsUsed {
		return
	}

	p.receivedTasksCount++
	if p.receivedTasksCount >= uint(p.workersCountLimit) {
		p.economyModeIsUsed = false
	}

	// Start an additional Worker.
	worker = NewWorker(p.receivedTasksCount, p)
	p.workers = append(p.workers, worker)
	err = worker.Start()
	if err != nil {
		return
	}
	p.workersWG.Add(1)
	return
}
