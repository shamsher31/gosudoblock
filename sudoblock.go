package sudoblock

import (
	"fmt"
	"github.com/shamsher31/goisroot"
	"github.com/shamsher31/gosymbol"
	"github.com/ttacon/chalk"
	"os"
)

func Is() {
	if root.Is() {
		fmt.Println(symbol.Error(), chalk.Red.Color("You are not allowed to run this app with root permissions"))
		os.Exit(77)
	}
}
