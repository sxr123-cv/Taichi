package log

import (
	"sync"
	"testing"
)

func TestLog_INFO(t *testing.T) {
	L := NewLog(nil, nil)
	wait := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i := i
		wait.Add(1)
		go func() {
			L.ERROR(i, "info")
			wait.Done()
		}()
	}
	wait.Wait()

}
