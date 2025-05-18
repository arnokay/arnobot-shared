package assert

import (
	"log"
	"log/slog"
)

func Assert(condition bool, msg string) {
	if !condition {
		slog.Error("assert:Assert: condition failed")
		log.Fatal(msg)
	}
}

func NoError(err error, msg string) {
	if err != nil {
		slog.Error("assert:NoError: error encountered", "error", err)
		log.Fatal(msg)
	}
}
