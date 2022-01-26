package convert

import (
	"fmt"
	"testing"
)

func TestMoreFloat2String(t *testing.T) {
	s := MoreFloat2String(float32(12.332453254363))
	fmt.Printf("val:%s type: %T\n", s, s)
	d := MoreFloat2String(23.434634643576435)
	fmt.Printf("val:%s type: %T\n", d, d)
}
