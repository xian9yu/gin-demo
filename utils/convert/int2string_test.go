package convert

import (
	"fmt"
	"testing"
)

func TestIntToString(t *testing.T) {

	s := IntToString(int(123))
	fmt.Printf("val:%s type: %T\n", s, s)
	d := IntToString(uint(234))
	fmt.Printf("val:%s type: %T\n", d, d)
	x := IntToString(int32(345))
	fmt.Printf("val:%s type: %T\n", x, x)
	a := IntToString(int64(456))
	fmt.Printf("val:%s type: %T\n", a, a)
}
