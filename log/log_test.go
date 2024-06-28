package log

import (
	"sync"
	"testing"
	"time"
)

func TestLog_INFO(t *testing.T) {
	L := NewLog(nil, nil)
	defer func() {
		L.Server.Close()
		time.Sleep(1 * time.Second)
	}()
	wait := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		i := 1
		go func() {
			L.ERROR(i, "info")
			wait.Done()
		}()
	}
	wait.Wait()
	return

}
