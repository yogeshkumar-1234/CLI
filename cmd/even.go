/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of your command",
	Long:  `Used to add even numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}
		str := fmt.Sprintf("The even addition of %s is %d", args, evenSum)
		fmt.Println(str)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)
}
