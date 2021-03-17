package util

import (
	"fmt"
	"os"
)

// GetAppPath returns current parogram path
func GetAppPath() (string, error) {
	c := os.Getenv("app.home")
	var err error
	if c == "" {
		c, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}
	return c, nil
}


// MakeDir makes dir given path.
func MakeDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("cannot create folder %s \n", path)
	}

}
