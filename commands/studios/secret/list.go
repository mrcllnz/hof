package secret

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listLong = `List your Studios secrets`

var ListCmd = &cobra.Command{

	Use: "list",

	Short: "List your secrets",

	Long: listLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("hof studios secret list")

	},
}
