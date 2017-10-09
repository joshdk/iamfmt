// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package iam

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func NewFromFile(name string) (*Policy, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return NewFromReader(file)
}

func NewFromString(s string) (*Policy, error) {
	src := strings.NewReader(s)

	return NewFromReader(src)
}

func NewFromReader(src io.Reader) (*Policy, error) {
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return NewFromBytes(data)
}

func NewFromBytes(buf []byte) (*Policy, error) {

	var policy Policy

	if err := json.Unmarshal(buf, &policy); err != nil {
		return nil, err
	}

	return &policy, nil
}
