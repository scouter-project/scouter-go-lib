package util

import (
	"fmt"
	"testing"
)

func TestOSUtil(t *testing.T) {
	p, _ := GetAppPath()
	fmt.Printf("%s", p)
}
