package cmd

import (
	"github.com/zoltan-magyar/golang-cli-template/internal/handlers"

	"github.com/spf13/cobra"
)

// goodbyeCmd represents the goodbye command
var goodbyeCmd = &cobra.Command{
	Use:   "goodbye",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: handlers.GoodbyeHandler,
}

func init() {
	rootCmd.AddCommand(goodbyeCmd)
}
