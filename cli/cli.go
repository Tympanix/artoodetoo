package cli

import (
	"fmt"
	"os"

	"github.com/Tympanix/automato/types"
	"github.com/mkideal/cli"
)

var main func(types.AppArgs)

// Run the cli application
func Run(fn func(types.AppArgs)) {
	main = fn

	prog := cli.Root(appCmd, cli.Tree(helpCmd), cli.Tree(addUserCmd), cli.Tree(genSecretCmd))

	if err := prog.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
