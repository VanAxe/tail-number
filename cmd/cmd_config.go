package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
  Use:   "config",
  Short: "Print out the interpreted configuration",
  Long: `Sometimes, you want to test out if what you configured is actually
  interpreted properly. Use this command to read out the converged and
  interpreted configuration settings.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("config called")
    fmt.Printf("%+v\n")
  },
}

func init() {
  RootCmd.AddCommand(configCmd)
}
