// If you have something like goplus in atom, or another auto-test-runner, something like this which has the potential
// to deadlock or leave behind artifacts on a failure can be incredibly annoying. So I broke it out into it's own
// little test utility as a smoke test to verify nothing can grab the lock when it's already held.
package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/rapid7/komand-plugin-sdk-go2/cache"
)

// Counter is a canary struct to alert us to issues
// It's designed without any locks, because if the cache lock works, it shouldn't need any
// It should panic / fatal under any potential issues, but since race conditions are race conditions
// you can never gaurantee when or if they will happen, so we need to run this a bunch of times
// and even then, nothing is perfect
// "You can't always prove something is threadsafe, you can only prove that it is not yet found to be thread-un-safe" : me
type counter struct {
	sync.Mutex
	count  int
	holder int
}

func (c *counter) Grab(grID int) {
	if c.count > 0 {
		log.Fatalf("%d tried to grab, but %d already held it\n", grID, c.holder)
	}
	c.count++
	c.holder = grID
}

func (c *counter) Release(grID int) {
	c.count--
	c.holder = -1
}

func main() {
	// Clean up last run
	os.Remove("/var/cache/lock/locker.lock")

	wg := sync.WaitGroup{}
	numWorkers := 10
	wg.Add(numWorkers)
	c := &counter{}
	// Fire off a bunch of go routines to contend for the lock
	for i := 0; i < numWorkers; i++ {
		go func(grID int) {
			timesRan := 0
			for timesRan < 500 {
				// Try to lock the cache`
				ok, err := cache.LockCacheFile("locker.lock")
				// If there is an error, it fails hard - shouldn't be an error
				if err != nil {
					log.Fatalf("%d\t:\t\t%s", grID, err)
				}
				// If we got the lock, then pretend to do some work and unlock it
				if ok {
					// Grabbing and Releasing the counter will be our "work"
					// If 2 things enter work block, the Grab call will log fatal if it's a proper race
					// They could still sneak in between eachother, but if the file lock is working properly, that can't happen
					// until AFTER cache.Unlock is called
					// What can I say, race conditions are hard to induce in a vacuum
					c.Grab(grID)
					c.Release(grID)
					dur := time.Millisecond * 1
					cache.UnlockCacheFile("locker.lock", &dur)
				}
				timesRan++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	// The program will exit 0 here if nothing went wrong
}
