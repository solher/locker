package lock

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

func newKey(values ...interface{}) string {
	var b bytes.Buffer
	for _, v := range values {
		b.WriteString(fmt.Sprintf("%v", v))
	}
	return b.String()
}

func TestNewEntityLocker(t *testing.T) {
	el := New()
	k1 := newKey("abc", 456)
	k2 := newKey("abcd", 456)

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		el.Lock(k1)
		fmt.Println(k1, 1)
		time.Sleep(time.Second * 4)
		fmt.Println("trying to unlock now")
		el.Unlock(k1)
	}()
	go func() {
		defer wg.Done()
		el.Lock(k1)
		fmt.Println(k1, 2)
		time.Sleep(time.Second * 4)
		el.Unlock(k1)
	}()
	go func() {
		defer wg.Done()
		el.Lock(k1)
		fmt.Println(k1, 3)
		time.Sleep(time.Second * 4)
		el.Unlock(k1)
	}()

	go func() {
		defer wg.Done()
		el.Lock(k1)
		fmt.Println(k2, 1)
		time.Sleep(time.Second * 4)
		el.Unlock(k1)
	}()

	wg.Wait()
}
