// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package iam

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func testPolicies() ([]string, error) {

	filenames := []string{}

	if err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err == nil && strings.HasSuffix(path, ".json") {
			filenames = append(filenames, path)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	sort.Strings(filenames)

	return filenames, nil
}

func TestParse(t *testing.T) {
	filenames, err := testPolicies()

	require.Nil(t, err)

	for index, path := range filenames {

		name := fmt.Sprintf("case #%d - %s", index, path)

		t.Run(name, func(t *testing.T) {
			_, err := NewFromFile(path)

			require.Nil(t, err)
		})

	}
}
