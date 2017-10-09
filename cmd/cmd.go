// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package cmd

import (
	"fmt"

	"github.com/palantir/pkg/cli"
)

func Cmd() *cli.App {

	app := cli.NewApp()

	app.Name = "iamfmt"
	app.Description = "AWS IAM policy document formatter"
	app.Version = Version()

	app.ErrorHandler = func(ctx cli.Context, err error) int {
		fmt.Printf("%s: %s\n", app.Name, err.Error())
		return 1
	}

	app.Action = func(ctx cli.Context) error {
		return nil
	}

	return app
}
