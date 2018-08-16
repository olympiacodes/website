package main

import (
	"io/ioutil"
)

//go:generate go run codegen/assets_generate.go

// Asset returns the bytes of a specific asset
func Asset(path string) ([]byte, error) {
	r, err := Assets.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}
