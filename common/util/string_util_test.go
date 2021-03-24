package util

import (
	"fmt"
	"testing"
)

func TestSnakeCase(t *testing.T) {
	s := "Test Case"
	s = SnakeCase(s)
	fmt.Println(s)
}
