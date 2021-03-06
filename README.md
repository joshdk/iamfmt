[![License](https://img.shields.io/github/license/joshdk/iamfmt.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/iamfmt?status.svg)](https://godoc.org/github.com/joshdk/iamfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/iamfmt)](https://goreportcard.com/report/github.com/joshdk/iamfmt)
[![CircleCI](https://circleci.com/gh/joshdk/iamfmt.svg?&style=shield)](https://circleci.com/gh/joshdk/iamfmt/tree/master)

# IAMFmt

📝 AWS IAM policy document formatter

## Goals 

### Data Correctness

Tooling should never change the functional effect of a policy.

### Git diff simplicity

If your project requires checking in a raw IAM policy, this tool should produce minimal diffs across commits.

### Idempotency

Running the tool multiple times should produce identical output, and running the library constituents should produce identical data structures.

### Formatting consistency

A policy should have a well-defined serialization ordering.

### Explicit over implicit

Implicit policy components should be made explicit.

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.