package cmd

import (
	"github.com/dsrosen6/addigy-command/addigy"
	"github.com/spf13/cobra"
)

var softwareId string

var installCmd = &cobra.Command{
	Use:   "install [id]",
	Short: "Install a software item from Addigy",
	Long: `
Starts an Addigy software item installation. Must be run as root (sudo).
To get the software item ID, navigate to the edit page for the software item in the Addigy web console.
In the URL, the ID is everything after "edit/".

Example Use Cases:
- You want to install a software item from the command line instead of from the Addigy web console.
- For installation from within another script or automation.
`,
	Args:             cobra.ExactArgs(1),
	PersistentPreRun: checkRoot,
	Run: func(cmd *cobra.Command, args []string) {
		softwareId = args[0]
		if spinner {
			if err := addigy.PolicierInstall(softwareId, true); err != nil {
				cmd.PrintErrln("Error installing software: ", err)
			}
			return
		}

		if err := addigy.PolicierInstall(softwareId, false); err != nil {
			cmd.PrintErrln("Error installing software: ", err)
		}
		return
	},
	DisableFlagsInUseLine: true,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"[id] - the software item ID to install"}, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&spinner, "spinner", "s", false, "install the software with a spinner instead of full output")
}
