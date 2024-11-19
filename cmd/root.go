package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.1.0"

var verFlag bool

var rootCmd = &cobra.Command{
	Use:   "addigy",
	Short: "A CLI tool for Addigy",
	Long:  `Addigy CLI is a command line tool to run Addigy-related commands from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verFlag {
			fmt.Println(version)
			return
		} else {
			_ = cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&verFlag, "version", "v", false, "get the current installed CLI version")
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func checkRoot(cmd *cobra.Command, args []string) {
	if os.Geteuid() != 0 {
		fmt.Printf("addigy %s must be run as root (sudo)\n", cmd.Name())
		os.Exit(1)
	}
}
