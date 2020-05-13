package util

import "io/ioutil"

//ReadFile returns file contents to string
func ReadFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(data)
}
