//go:build tools

// This file used for attaching tools dependencies to the project
package zksync

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/vektra/mockery/cmd/mockery"
)
