package convert

import (
	"fmt"
	"testing"
)

func TestMoreInt2String(t *testing.T) {

	s := MoreInt2String(int(123))
	fmt.Printf("val:%s type: %T\n", s, s)
	d := MoreInt2String(uint(234))
	fmt.Printf("val:%s type: %T\n", d, d)
	x := MoreInt2String(int32(345))
	fmt.Printf("val:%s type: %T\n", x, x)
	a := MoreInt2String(int64(456))
	fmt.Printf("val:%s type: %T\n", a, a)
}
