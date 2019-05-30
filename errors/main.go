package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConf() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settting.xml"))
	return config, errors.WithMessage(err, "could not read config")
}
func main() {
	_, err := ReadConf()
	if err != nil {
		fmt.Printf("original error:%T\n%v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}
