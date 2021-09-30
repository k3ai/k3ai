package log

import (
	"os"
)

func CheckErrors(err error) error {
	if err != nil {
		Error(err)
		os.Exit(0)
	}
	return nil
}