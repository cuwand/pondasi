package hmac

import (
	"fmt"
	"testing"
)

func TestHashed(t *testing.T) {
	fmt.Println(Hashed("mysecret", "data"))
}
