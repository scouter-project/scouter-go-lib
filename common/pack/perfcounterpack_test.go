package pack

import (
	"fmt"
	"testing"

	"github.com/scouter-project/scouter-go-lib/common/netdata"
)

func TestPerfCounterPack(t *testing.T) {
	pack := NewPerfCounterPack()
	pack.Put("fvalue", 1.23)
	pack.Put("ivalue", 123)
	pack.Put("tvalue", " test value")
	pack.Put("bvalue", false)
	fmt.Printf("%s", pack.ToString())
}

func TestPerfCounterPack2(t *testing.T) {
	pack := NewPerfCounterPack()
	pack.Put("fvalue", 1.23)
	pack.Put("ivalue", 123)
	pack.Put("tvalue", " test value2")
	pack.Put("bvalue", false)
	out := netdata.NewDataOutputX()
	pack.Write(out)
	in := netdata.NewDataInputX(out.Bytes())
	pack2 := NewPerfCounterPack()
	pack2.Read(in)
	fmt.Printf("%s", pack2.ToString())
}
