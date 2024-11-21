package codazap_test

import (
	"github.com/marnixbouhuis/coda"
	codazap "github.com/marnixbouhuis/coda-zap"
	"go.uber.org/zap"
)

func Example() {
	// Initialize zap logger
	l, _ := zap.NewDevelopment()

	codaLogger := codazap.NewLogger(l.Named("coda"))
	_ = coda.NewShutdown(
		// Register logger
		coda.WithShutdownLogger(codaLogger),
	)
}
