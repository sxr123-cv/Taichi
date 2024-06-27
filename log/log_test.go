package log

import "testing"

func TestLog_INFO(t *testing.T) {
	L := NewLog(nil, nil)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			L.INFO(i, "info")
		}()
	}

}

func TestLog_INFO2(t *testing.T) {

}

func TestLog_INFO3(t *testing.T) {

}
