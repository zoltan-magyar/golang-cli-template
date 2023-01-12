package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GreetHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to the CLI client!")
}
