package cmd

import (
	"github.com/zoltan-magyar/golang-cli-template/internal/handlers"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,

	Run: handlers.GreetHandler,
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
