package convert

import (
	"fmt"
	"testing"
)

func TestFloatToString(t *testing.T) {
	s := FloatToString(float32(12.332453254363))
	fmt.Printf("val:%s type: %T\n", s, s)
	d := FloatToString(23.434634643576435)
	fmt.Printf("val:%s type: %T\n", d, d)
}
