package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GoodbyeHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Goodbye!")
}
