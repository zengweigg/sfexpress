package tool

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	b := PKCS7Padding(566)
	fmt.Println(b)
}
