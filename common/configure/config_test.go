package configure

import (
	"os"
	"testing"
	"time"
)

func TestConfigure_Get(t *testing.T) {
	//configure support system env value
	os.Setenv("key1","value1")
	os.Setenv("key2", "value2")
	// configure read config file whenever config file changes
	conf := NewConfigure("/etc/my.conf")
	//confiugre support manually add config value
	conf.Put("key3","value3")
	conf.Start()
	for true {
		time.Sleep(10 * time.Second)
	}
}