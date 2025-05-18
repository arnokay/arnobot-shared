package assert

import (
	"log/slog"
)

func Assert(condition bool, msg string) {
	if !condition {
		slog.Error("assert:Assert: condition failed", "msg", msg)
		panic(msg)
	}
}

func NoError(err error, msg string) {
	if err != nil {
		slog.Error("assert:NoError: error encountered", "error", err, "msg", msg)
    panic(msg)
	}
}

