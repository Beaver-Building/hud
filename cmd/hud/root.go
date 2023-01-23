package hud

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boardgamemaster",
	Short: "bgm - a simple CLI to transform and inspect player data",
	Long: `bgm is a super fancy CLI (kidding)
   
One can use bgm to modify or inspect player data straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
