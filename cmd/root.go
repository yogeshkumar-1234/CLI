/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("Starting CLI") },
}

func Execute() {
	fmt.Println("Execute root")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// fmt.Println("Root dir")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// init Constructor
// Execute -> Run
