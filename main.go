package main

import "git.1vh.de/maximilian.pachl/bowler/bowlerfile"
import "os"
import "fmt"
import "path/filepath"


// ----------------------------------------------------------------------------------
//  Konstanten
// ----------------------------------------------------------------------------------

/** Name des Buildfiles. */
const BUILD_FILE_NAME = "Bowlerfile"


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func main() {
	path, err := filepath.Abs(BUILD_FILE_NAME)
	if (err == nil) {
		fmt.Println("Using buildfile " + path)
		fmt.Println()
	}

	// Load the Bowlerfile
	buildFile, err := bowlerfile.Load(BUILD_FILE_NAME)
	if (err != nil) {
		fmt.Printf("Could not open Bowlerfile: %s\n", err)
		os.Exit(-3)
	}

	if (len(os.Args) >= 2) {
		// the user requests us to build the project
		if (os.Args[1] == "build") {
			fmt.Println("Executing task 'build':")
			build(buildFile)	

		// clean the project directory
		} else if (os.Args[1] == "clean") {
			fmt.Println("Executing task 'clean':")
			BeginStepMessage("Cleaning project root")
			err = clean(buildFile)
			EndStepMessage(err)

		// i don't know what the user want form me :O
		} else {
			fmt.Println("ERROR: invalid command: " + os.Args[1])
		}

	// the subcommand was not provided -> error	
	} else {
		fmt.Println("ERROR: missing command")
	}
}