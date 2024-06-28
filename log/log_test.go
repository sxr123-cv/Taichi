package log

import (
	"testing"
)

func TestLog_INFO(t *testing.T) {
	L := NewLog(nil, nil)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			L.ERROR(i, "info")
		}()
	}
	for {

	}

}

func TestLog_INFO2(t *testing.T) {

}

func TestLog_INFO3(t *testing.T) {

}
