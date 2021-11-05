package files

import (
	"io/ioutil"
)

func ReadFile(filename string) ([]byte, error) {
	s, _ := ioutil.ReadFile(filename)
	return s, nil
}
