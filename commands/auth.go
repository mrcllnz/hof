package commands

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/commands/auth"
)

var authLong = `authentication subcommands`

var AuthCmd = &cobra.Command{

	Use: "auth",

	Short: "authentication subcommands",

	Long: authLong,
}

func init() {
	AuthCmd.AddCommand(auth.LoginCmd)
}
