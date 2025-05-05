package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString("name")
		if name == "" {
			name = "World"
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	helloCmd.Flags().String("name", "", "Name to greet")
	viper.BindPFlag("name", helloCmd.Flags().Lookup("name"))
	rootCmd.AddCommand(helloCmd)
}
