package addigy

import (
	"fmt"
	"os"
)

const (
	addigyFolder = "/Library/Addigy/"
	goAgent      = addigyFolder + "go-agent"
	statusPath   = addigyFolder + "ansible/status.json"
)

type command struct {
	mainCommand string
	args        []string
}

var policierRunCmd = command{
	mainCommand: goAgent,
	args:        []string{"policier", "run"},
}

// PolicierRun starts the Addigy policier, and just lets you know if it started.
func PolicierRun(mode string) error {
	switch mode {
	case "verbose":
		return commandWithOutput(policierRunCmd)

	case "spinner":
		action := func() {
			err := commandWithoutOutput(policierRunCmd)
			if err != nil {
				fmt.Println("Error running Addigy Policy: ", err)
				return
			}
		}
		RunWithSpinner(" Running the Addigy Policy...this may take a few minutes.", action)
		fmt.Println("Addigy Policy run complete.")
		return nil

	default:
		c := command{
			mainCommand: "launchctl",
			args:        []string{"start", "com.addigy.policier"},
		}

		err := commandWithoutOutput(c)
		if err != nil {
			return fmt.Errorf("could not run Addigy Policy: %w", err)
		}
		fmt.Println("Addigy Policy run started.")
		return nil
	}
}

func ResetPolicyProgress() error {
	err := os.Remove(statusPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No policy progress to reset.")
			return nil
		}
		return fmt.Errorf("could not reset policy progress: %w", err)
	}
	fmt.Println("Successfully reset policy progress.")
	return nil
}

func PolicierInstall(item string, useSpinner bool) error {
	c := command{
		mainCommand: goAgent,
		args:        []string{"policier", "install", item},
	}

	if useSpinner {
		action := func() {
			err := commandWithoutOutput(c)
			if err != nil {
				fmt.Println("Error installing software: ", err)
				return
			}
		}
		RunWithSpinner(" Installing software...", action)
		fmt.Println("Software installed.")
		return nil
	}

	return commandWithOutput(c)

}
