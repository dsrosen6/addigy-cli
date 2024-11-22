# addigy-cli
Addigy CLI is a command line tool to run Addigy-related commands from the terminal.

## Run
___
Starts an Addigy policy run. Must be run as root (sudo).  
Without flags, it will only let you know the policy run has started.

### Example Use Cases
- You want to start a policy run without waiting for the next scheduled run.
- You want confirmation that the policy run has started and finished (with verbose or spinner flags).
- You just deployed a new software item and want to make sure it runs ASAP on this specific Mac.

### Usage
`addigy run [flags]`

### Flags
`-s, --spinner`: run the policy with a spinner. Cannot be used with --verbose.  
`-v, --verbose`: run the policy with full verbose output. Cannot be used with --spinner.

## Reset
___
Resets all Addigy policy progress. Must be run as root (sudo).

### Example Use Cases
- You want Addigy to re-run any "run once" software items (those without condition scripts).
- There are software items that are stuck in a pending state which won't change after a policy run.

### Usage 
`addigy reset [flags]`

### Flags 
`-h, --help`: help for reset

## Install
___
Starts an Addigy software item installation. Must be run as root (sudo).  
*To get the software item ID, navigate to the edit page for the software item in the Addigy web console.
In the URL, the ID is everything after "edit/".*

### Example Use Cases
- You want to install a software item from the command line instead of from the Addigy web console.
- For installation from within another script or automation.

### Usage
`addigy install [id]`

### Flags  
`-h, --help`: help for install  
`-s, --spinner`: install the software with a spinner instead of full output