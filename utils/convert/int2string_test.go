package convert

import (
	"fmt"
	"testing"
)

func TestIntToString(t *testing.T) {
	a := IntToString(uint(111))
	fmt.Printf("val:%s type: %T\n", a, a)
	b := IntToString(uint16(222))
	fmt.Printf("val:%s type: %T\n", b, b)
	c := IntToString(uint32(333))
	fmt.Printf("val:%s type: %T\n", c, c)
	d := IntToString(uint64(444))
	fmt.Printf("val:%s type: %T\n", d, d)

	e := IntToString(int(123))
	fmt.Printf("val:%s type: %T\n", e, e)
	f := IntToString(uint(234))
	fmt.Printf("val:%s type: %T\n", f, f)
	g := IntToString(int32(345))
	fmt.Printf("val:%s type: %T\n", g, g)
	h := IntToString(int64(456))
	fmt.Printf("val:%s type: %T\n", h, h)
}
