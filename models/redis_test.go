package models

import (
	"fmt"
	"testing"
)

func TestStrExists(t *testing.T) {
	val, err := StrExists("asddsa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
