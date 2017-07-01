package main

import (
	"github.com/goph/emperror"
	"github.com/goph/log"
	"github.com/goph/stdlib/ext"
)

// newErrorHandler creates a new Emperror error handler.
func newErrorHandler(config *configuration, logger log.Logger) (emperror.Handler, ext.Closer) {
	return emperror.NewLogHandler(logger), ext.NoopCloser
}
