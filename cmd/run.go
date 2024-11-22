package cmd

import (
	"github.com/dsrosen6/addigy-command/addigy"
	"github.com/spf13/cobra"
	"os"
)

var (
	verbose bool
	spinner bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts a policy run",
	Long: `
Starts an Addigy policy run. Must be run as root (sudo).
Without flags, it will only let you know the policy run has started.

Example Use Cases:
- You want to start a policy run without waiting for the next scheduled run.
- You want confirmation that the policy run has started and finished (with verbose or spinner flags).
- You just deployed a new software item and want to make sure it runs ASAP on this specific Mac.
`,
	PersistentPreRun: checkRoot,
	Run: func(cmd *cobra.Command, args []string) {
		if spinner && verbose {
			cmd.PrintErrln("Cannot use both --spinner and --verbose flags together.")
			os.Exit(1)
		}

		if spinner {
			if err := addigy.PolicierRun("spinner"); err != nil {
				cmd.PrintErrln("Error: ", err)
				os.Exit(1)
			}
			return
		}

		if verbose {
			if err := addigy.PolicierRun("verbose"); err != nil {
				cmd.PrintErrln("Error: ", err)
				os.Exit(1)
			}
			return
		}

		if err := addigy.PolicierRun("default"); err != nil {
			cmd.PrintErrln("Error: ", err)
			os.Exit(1)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "run the policy with full verbose output. Cannot be used with --spinner.")
	runCmd.Flags().BoolVarP(&spinner, "spinner", "s", false, "run the policy with a spinner. Cannot be used with --verbose.")
}
