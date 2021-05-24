package smash

import (
	"bufio"
	"io"
	"sync"
)

type word string
var wg sync.WaitGroup
var mu sync.Mutex

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)

	// scan the reader into the scanner
	scanner := bufio.NewScanner(r)
	// split by spaces
	scanner.Split(bufio.ScanWords)
	// for each word; concurrently call the smasher function 
	for scanner.Scan() {
		val := smasher(word(scanner.Text()))
		wg.Add(1)
		go UpdateVal(val, m)
	}
	wg.Wait()
	return m
}

func UpdateVal(keyVal uint32, M map[uint32]uint) {
	//lock during critical section; then uodate wait group 
	mu.Lock()
	M[keyVal]++
	mu.Unlock()
	wg.Done()
}
