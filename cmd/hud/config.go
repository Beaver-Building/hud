package hud

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		/* LOAD MECHANISM
		** get pool
		** get victory conditions
		** get players
		 */
		/* FUNCTIONAL COMPONENTS
		** load config from *fandom* server TBD
		** connect to cosmos db emulator https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/samples-go
		** connect to azurite for queue https://learn.microsoft.com/en-us/azure/storage/common/storage-use-azurite?tabs=visual-studio
		 */

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
