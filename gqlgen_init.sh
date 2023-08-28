#!bin/bash

echo "//go:build tools
// +build tools

package tools

import (
	_ \"github.com/99designs/gqlgen\"
)" > tools.go

go run github.com/99designs/gqlgen init

go mod tidy