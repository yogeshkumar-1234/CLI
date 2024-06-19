/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func addInt(args []string) {
	var sum int
	for _, val := range args {
		itemp, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	str := fmt.Sprintf("Addition of numbers %s is %d", args, sum)
	fmt.Println(str)
}

func addFloat(args []string) {
	var sum float64
	for _, val := range args {
		itemp, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	str := fmt.Sprintf("Addition of numbers %s is %f", args, sum)
	fmt.Println(str)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Used to add Numbers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating numbers")
	// --float or -f
}
