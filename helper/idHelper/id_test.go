package idHelper

import (
	"fmt"
	"testing"
)

func TestId(t *testing.T) {
	for i := 0; i < 100000; i++ {
		//UUID() // 0.15s, 0.15s, 0.14s
		//UUIDClean() // 0.16s, 0.16s, 0.16s
		ULID() // 0.03s, 0.02s, 0.02s
	}

	fmt.Println("OK SIP!")
}
