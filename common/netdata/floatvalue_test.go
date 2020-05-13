package netdata

import (
	"fmt"
	"testing"
)

func TestFloatValue(t *testing.T) {
	out := NewDataOutputX()
	fvalue := NewFloatValue(124.234)
	fvalue.Write(out)
	in := NewDataInputX(out.Bytes())
	fmt.Printf("value : %f\n", in.ReadFloat32())
	fmt.Printf("value %s\n", fvalue.ToString())
}
