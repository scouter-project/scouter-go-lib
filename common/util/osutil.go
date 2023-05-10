package util

import (
	"flag"
	"fmt"
	"os"
)

// GetAppPath returns current program path
func GetAppPath() (string, error) {
	c := os.Getenv("app.home")
	if c == "" {
		if flag.Lookup("app.home") == nil {
			flag.StringVar(&c, "app.home", "", "application home")
			//flag.Parse()
		} else {
			c = flag.Lookup("app.home").Value.String()
		}
	}
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
