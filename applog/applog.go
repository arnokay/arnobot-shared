package applog

import (
	"os"
)

var l Logger = NewSlogLogger(os.Stdout, "", 0, nil)

func SetDefault(logger Logger) {
  l = logger
}

func Default() Logger {
  return l
}

func NewWithAttrs(attrs ...any) Logger {
  return Default().With(attrs...)
}

func NewServiceLogger(serviceName string, attributes ...any) Logger {
	attributes = append(attributes, LoggerServiceKey, serviceName)

  return NewWithAttrs(attributes...)
}
