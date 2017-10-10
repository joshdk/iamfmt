// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package iam

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//func TestUnmarshall(t *testing.T) {
//
//	tests := []struct {
//		name     string
//		json     string
//		expected []string
//	}{
//		{
//			name:     "null",
//			json:     `null`,
//			expected: []string{},
//		},
//		{
//			name:     "single string",
//			json:     `"alice"`,
//			expected: []string{"alice"},
//		},
//		{
//			name:     "empty list",
//			json:     `[]`,
//			expected: []string{},
//		},
//		{
//			name:     "single string list",
//			json:     `["alice"]`,
//			expected: []string{"alice"},
//		},
//		{
//			name:     "multiple string list",
//			json:     `["alice", "bob", "carol"]`,
//			expected: []string{"alice", "bob", "carol"},
//		},
//	}
//
//	for index, test := range tests {
//		name := fmt.Sprintf("case #%d - %s", index, test.name)
//
//		t.Run(name, func(t *testing.T) {
//
//			var actual stringSlice
//
//			err := json.Unmarshal([]byte(test.json), &actual)
//
//			require.Nil(t, err)
//
//			assert.Equal(t, test.expected, []string(actual))
//		})
//	}
//}

func TestUnmarshallClause(t *testing.T) {

	tests := []struct {
		name     string
		json     string
		expected []string
	}{
		{
			name: "single string",
			json: `
			{
				"Action": "alice"
			}`,
			expected: []string{"alice"},
		},
		{
			name: "empty list",
			json: `
			{
				"Action": []
			}`,
			expected: []string{},
		},
		{
			name: "single string list",
			json: `
			{
				"Action": [
					"alice"
				]
			}`,
			expected: []string{"alice"},
		},
		{
			name: "multiple string list",
			json: `
			{
				"Action": [
					"alice",
					"bob",
					"carol"
				]
			}`,
			expected: []string{"alice", "bob", "carol"},
		},
	}

	for index, test := range tests {
		name := fmt.Sprintf("case #%d - %s", index, test.name)

		t.Run(name, func(t *testing.T) {

			var actual Statement

			err := json.Unmarshal([]byte(test.json), &actual)

			require.Nil(t, err)

			assert.Equal(t, test.expected, []string(actual.Actions))
		})
	}
}
