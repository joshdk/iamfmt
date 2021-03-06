// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package iam

import (
	"encoding/json"
)

type Policy struct {
	Version    string     `json:"Version"`
	Statements Statements `json:"Statement"`
}

type Statement struct {
	Effect     string    `json:"Effect"`
	Actions    Actions   `json:"Action"`
	Resources  Resources `json:"Resource"`
	Principals Principal `json:"Principal"`
}

type Statements []Statement

func (statements *Statements) UnmarshalJSON(body []byte) (err error) {

	var (
		stmts []Statement
		stmt  Statement
	)

	// If the value is a list of Statement, then return that list of Statement
	if err = json.Unmarshal(body, &stmts); err == nil {
		*statements = stmts
		return
	}

	// If the value is a single Statement, then return a list containing only that Statement
	if err = json.Unmarshal(body, &stmt); err == nil {
		*statements = []Statement{stmt}
		return
	}

	return
}

type Principal struct {
	Service []string `json:"Service"`
}

type Actions []string

func (actions *Actions) UnmarshalJSON(body []byte) (err error) {
	*actions, err = fromSingleOrList(body)
	return err
}

type Resources []string

func (resources *Resources) UnmarshalJSON(body []byte) (err error) {
	*resources, err = fromSingleOrList(body)
	return err
}

func fromSingleOrList(body []byte) ([]string, error) {
	var (
		items []string
		item  string
		err   error
	)

	// If the value is a list of strings, then return that list of strings
	if err = json.Unmarshal(body, &items); err == nil {
		return items, nil
	}

	// If the value is a single string, then return a list containing only that string
	if err = json.Unmarshal(body, &item); err == nil {
		return []string{item}, nil
	}

	return nil, err
}
