# Coda-zap - Zap log adapter for Coda
[![Go Reference](https://pkg.go.dev/badge/github.com/marnixbouhuis/coda-zap.svg)](https://pkg.go.dev/github.com/marnixbouhuis/coda-zap)
[![CI/CD Pipeline](https://github.com/MarnixBouhuis/coda-zap/actions/workflows/cicd.yaml/badge.svg)](https://github.com/marnixbouhuis/coda-zap/actions/workflows/cicd.yaml)

https://github.com/marnixbouhuis/coda

## Installation

```bash
go get github.com/marnixbouhuis/coda-zap
```

## Quick Start

```go
package main

import (
	"github.com/marnixbouhuis/coda"
	codazap "github.com/marnixbouhuis/coda-zap"
	"go.uber.org/zap"
)

func main() {
	// Initialize zap logger
	l, _ := zap.NewDevelopment()

	codaLogger := codazap.NewLogger(l.Named("coda"))
	_ = coda.NewShutdown(
		// Register logger
		coda.WithShutdownLogger(codaLogger),
	)
}
```
