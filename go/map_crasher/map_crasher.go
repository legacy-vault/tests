// map_crasher.go

package main

import (
	"sync"
)

var racer map[int]int

var race sync.RWMutex

func Read() {

	race.RLock()
	for k, v := range racer {
		_, _ = k, v
	}
	race.RUnlock()
}

func Write() {

	for i := 0; i < 1e8; i++ {
		race.Lock()
		racer[i/2] = i
		race.Unlock()
	}
}

func main() {
	racer = make(map[int]int)
	Write()
	go Write()
	Read()
}
