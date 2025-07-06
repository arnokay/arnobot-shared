package assert

import (
	"log"
	"os"

	"github.com/arnokay/arnobot-shared/applog"
)

func Assert(condition bool, msg string) {
	if !condition {
		applog.Default().Error("assert:Assert: condition failed")
		log.Fatal(msg)
	}
}

func NoError(err error, msg string) {
	if err != nil {
		applog.Default().Error("assert:NoError: error encountered", "message", msg, "error", err)
    os.Exit(1)
	}
}
