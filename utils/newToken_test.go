package utils

import (
	"fmt"
	"testing"
)

func TestNewToken(t *testing.T) {
	s := NewToken("xian9yu", "1")
	fmt.Println(s)
}
