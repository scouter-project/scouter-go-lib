package netdata

import (
	"fmt"
	"testing"
)

func TestDecimalValue(t *testing.T) {
	out := NewDataOutputX()
	dvalue := NewDecimalValue(1234)
	dvalue.Write(out)
	in := NewDataInputX(out.Bytes())
	result := in.ReadDecimal()
	fmt.Printf("value: %d\n", result)
}
