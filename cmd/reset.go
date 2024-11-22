package cmd

import (
	"github.com/dsrosen6/addigy-command/addigy"
	"github.com/spf13/cobra"
	"os"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets policy progress",
	Long: `
Resets all Addigy policy progress. Must be run as root (sudo).

Example Use Cases:
- You want Addigy to re-run any "run once" software items (those without condition scripts).
- There are software items that are stuck in a pending state which won't change after a policy run.
`,
	PersistentPreRun: checkRoot,
	Run: func(cmd *cobra.Command, args []string) {
		if err := addigy.ResetPolicyProgress(); err != nil {
			cmd.PrintErrln("Error: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
