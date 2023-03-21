package main

import (
	"io/ioutil"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	file := testFile
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic("failed to read file: " + file)
	}
	getFuncName(data)
}
