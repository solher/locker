package lock

import (
	"fmt"
	"sync"
)

type entityLock struct {
	sync.Mutex
	References int64
}

// EntityLocker contains locks in a map. Locks get aquired per map key.
type EntityLocker struct {
	locks   map[string]*entityLock
	mapLock *sync.Mutex
}

// New creates a new EntityLocker.
func New() *EntityLocker {
	return &EntityLocker{
		locks:   make(map[string]*entityLock),
		mapLock: new(sync.Mutex),
	}
}

// Lock aquires a lock for the given key.
func (lker *EntityLocker) Lock(key string) {
	lker.mapLock.Lock()
	lk, ok := lker.locks[key]
	if !ok {
		lk = new(entityLock)
		lker.locks[key] = lk
	}
	lk.References++
	lker.mapLock.Unlock()

	lk.Lock()
}

// Unlock releases the lock for the given key.
func (lker *EntityLocker) Unlock(key string) {
	lker.mapLock.Lock()
	lk, ok := lker.locks[key]
	if !ok {
		panic(fmt.Errorf("BUG: Lock for key '%s' not initialized", key))
	}
	lk.References--
	if lk.References == 0 {
		delete(lker.locks, key)
	}
	lker.mapLock.Unlock()

	lk.Unlock()
}
