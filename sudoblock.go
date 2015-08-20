package sudoblock

import (
	"fmt"
	"github.com/shamsher31/goisroot"
	"github.com/shiena/ansicolor"
	"os"
)

func Is() {
	if root.Is() {
		w := ansicolor.NewAnsiColorWriter(os.Stdout)
		message := "%sYou are not allowed to run this app with root permissions.%s\n"
		fmt.Fprintf(w, message, "\x1b[31m", "\x1b[0m")
		os.Exit(77)
	}
}
