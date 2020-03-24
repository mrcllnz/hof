package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var setLong = `set configuration values`

var SetCmd = &cobra.Command{

	Use: "set <name> <host> <account> [project]",

	Short: "set configuration values",

	Long: setLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]

		}

		if 1 >= len(args) {
			fmt.Println("missing required argument: 'Host'")
			cmd.Usage()
			os.Exit(1)
		}

		var host string

		if 1 < len(args) {

			host = args[1]

		}

		if 2 >= len(args) {
			fmt.Println("missing required argument: 'Account'")
			cmd.Usage()
			os.Exit(1)
		}

		var account string

		if 2 < len(args) {

			account = args[2]

		}

		var project string

		if 3 < len(args) {

			project = args[3]

		}

		// Default body

		fmt.Println("hof config set", name, host, account, project)

	},
}
